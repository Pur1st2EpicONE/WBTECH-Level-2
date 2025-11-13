// Package input parses command-line flags and arguments used by the
// wget application.
package input

import (
	"errors"
	"net/url"
	"strconv"
	"strings"

	"L2.16/internal/config"
	"github.com/spf13/pflag"
)

var (
	ErrMissingURL    = errors.New("missing URL")
	ErrInvalidDomain = errors.New("failed to parse domain")
)

// Flags holds parsed command-line flag values.
type Flags struct {
	Flagr bool
	Flagl uint
	lRaw  string
}

// Data aggregates parsed flags, the provided URL, domain and the loaded
// configuration.
type Data struct {
	Flags  Flags
	URL    string
	Domain string
	Config *config.Config
}

// ParseArgs reads and validates CLI flags and arguments. It returns a Data
// struct ready to be used by the rest of the program or an error when
// required arguments are missing or invalid.
func ParseArgs() (*Data, error) {

	input := new(Data)
	var err error

	pflag.BoolVarP(&input.Flags.Flagr, "recursive", "r", false, "specify recursive download")
	pflag.StringVarP(&input.Flags.lRaw, "level", "l", "0", "maximum recursion depth (inf or 0 for infinite)")

	pflag.Parse()

	args := pflag.Args()
	if len(args) == 0 {
		return nil, ErrMissingURL
	}

	input.URL = args[0]
	input.Domain, err = parseDomain(input.URL)
	if err != nil {
		return input, ErrInvalidDomain
	}

	if input.Flags.Flagr {
		lCheck(&input.Flags.Flagl, input.Flags.lRaw)
	} else {
		input.Flags.Flagl = 1
	}

	return input, nil

}

// lCheck parses recursion depth strings such as "inf", "0" or a number and
// sets the uint pointer accordingly.
func lCheck(flagl *uint, lString string) {

	if lString == "inf" || lString == "0" {
		return
	}

	recursionDepth, err := strconv.Atoi(lString)
	if err != nil {
		*flagl = 0
	} else {
		*flagl = uint(recursionDepth)
	}

}

// parseDomain extracts the host component from a raw URL string. It returns
// the host or an error if the URL cannot be parsed.
func parseDomain(rawURL string) (string, error) {

	if !strings.Contains(rawURL, "://") {
		rawURL = "http://" + rawURL
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return u.Host, nil

}
