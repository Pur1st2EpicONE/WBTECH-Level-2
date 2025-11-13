package input

import (
	"os"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestParseArgs_MissingURL(t *testing.T) {

	defaultSet := pflag.CommandLine
	defer func() { pflag.CommandLine = defaultSet }()

	pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	os.Args = []string{"wget"}

	_, err := ParseArgs()
	assert.Error(t, err, "expected ErrMissingURL")
	assert.Equal(t, ErrMissingURL, err, "expected ErrMissingURL")

}

func Test_parseDomain(t *testing.T) {

	tests := []struct {
		in  string
		exp string
	}{
		{"aboba.ru", "aboba.ru"},
		{"http://input.org/aboba", "input.org"},
		{"https://qweqwe.com:8080/foo", "qweqwe.com:8080"},
	}

	for _, tc := range tests {
		got, err := parseDomain(tc.in)
		assert.NoError(t, err, "parseDomain(%q) returned error", tc.in)
		assert.Equal(t, tc.exp, got, "parseDomain(%q) = %q; want %q", tc.in, got, tc.exp)
	}

}

func Test_lCheck(t *testing.T) {

	var v uint

	lCheck(&v, "inf")
	assert.Equal(t, uint(0), v, "expected 0 for 'inf'")

	lCheck(&v, "0")
	assert.Equal(t, uint(0), v, "expected 0 for '0'")

	lCheck(&v, "5")
	assert.Equal(t, uint(5), v, "expected 5 for '5'")

	lCheck(&v, "not-a-number")
	assert.Equal(t, uint(0), v, "expected 0 for invalid number")

}
