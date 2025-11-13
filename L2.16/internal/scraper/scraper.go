// Package scraper defines the Scraper abstraction and provides a factory to
// create an implementation backed by the colly-based scraper.
package scraper

import (
	"time"

	"L2.16/internal/input"
	"L2.16/internal/scraper/colly"
)

// Scraper represents the minimal scraping functionality used by the app.
// Implementations must provide a Scrape method and a GetStats helper.
type Scraper interface {
	// Scrape starts scraping the given URL according to the scraper's configuration.
	Scrape(URL string)
	// GetStats returns statistics about the scraping session,
	// including total files, total size, and elapsed time.
	GetStats() (int, int64, time.Duration)
}

// NewScraper returns a Scraper implementation using the colly-based
// scraper. The provided input.Data is used to configure the scraper.
func NewScraper(data *input.Data) Scraper {
	return colly.NewCollyScraper(data)
}
