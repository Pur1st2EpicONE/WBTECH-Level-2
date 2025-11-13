// Package wget contains the main command logic for the wget-like application.
// It handles argument parsing, configuration loading, scraper creation,
// execution, and printing of final statistics.
package wget

import (
	"errors"
	"os"

	"L2.16/internal/config"
	"L2.16/internal/input"
	"L2.16/internal/scraper"
	"L2.16/internal/utils"
)

// Wget is the entry point used by cmd/wget/main.go. It orchestrates argument
// parsing, configuration loading (with a fallback to defaults) and runs the
// configured scraper. Errors are reported to stderr and on fatal conditions
// the process exits with a non-zero code.
func Wget() {

	data, err := input.ParseArgs()
	if err != nil {
		if errors.Is(err, input.ErrMissingURL) {
			utils.LogError("wget: missing URL", nil)
			utils.LogError("Usage: wget [OPTION]... [URL]...", nil)
		} else {
			utils.LogError("wget: fatal: failed to parse command-line arguments", err)
		}
		os.Exit(1)
	}

	data.Config, err = config.Load()
	if err != nil {
		switch {
		case errors.Is(err, config.ErrLoadConfig):
			utils.LogError("wget: failed to read config file", err)
		case errors.Is(err, config.ErrUnmarshalConfig):
			utils.LogError("wget: invalid format in config file", err)
		case errors.Is(err, config.ErrNoSelectors):
			utils.LogError("wget: no download targets specified in config.", nil)
		}
		utils.LogError("wget: falling back to default configuration\n", nil)
		data.Config = config.Default()
	}

	s := scraper.NewScraper(data)

	s.Scrape(data.URL)

	utils.FormatOutput(s.GetStats())

}
