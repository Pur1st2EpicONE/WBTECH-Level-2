// Package input handles parsing command-line arguments and flags for grep.
package input

import (
	"errors"
	"regexp"

	"L2.12/internal/buffer"
	"github.com/spf13/pflag"
)

// Errors from ParseArgs, checked in logFatal if returned.
var (
	ErrInvalidContext = errors.New("invalid context length argument")
	ErrNoArgs         = errors.New("no args provided")
)

// Flags stores command-line options that modify grep behavior.
type Flags struct {
	FlagA int  // -A, number of lines after match
	FlagB int  // -B, number of lines before match
	FlagC int  // -C, number of lines for both before and after context
	Flagc bool // -c, count matching lines
	Flagi bool // -i, ignore case
	Flagv bool // -v, invert match
	FlagF bool // -F, fixed string match
	Flagn bool // -n, print line numbers
}

// Data represents the complete context of a grep operation,
// including parsed flags, input files, regex pattern, and printing state.
type Data struct {
	Flags                 Flags          // Parsed command-line flags
	MatchFound            bool           // True if any line matched the pattern
	CurrLineNumber        int            // Current line number being processed
	LastPrintedLineNumber int            // Line number of the last printed line
	LinesAfter            int            // Number of trailing lines to print after a match
	LinesBefore           *buffer.Buffer // Buffer storing preceding lines for context
	Pattern               *regexp.Regexp // Compiled regex pattern for matching
	FileNames             []string       // List of input file names
	CurrentFile           string         // Name of the currently processed file
}

// ParseArgs parses command-line arguments into a Data structure.
// It validates context flags and compiles the search pattern.
func ParseArgs() (*Data, error) {

	input := new(Data)

	pflag.IntVarP(&input.Flags.FlagA, "after-context", "A", 0, "print NUM lines of trailing context")
	pflag.IntVarP(&input.Flags.FlagB, "before-context", "B", 0, "print NUM lines of leading context")
	pflag.IntVarP(&input.Flags.FlagC, "context", "C", 0, "print NUM lines of output context")
	pflag.BoolVarP(&input.Flags.Flagc, "count", "c", false, "print only a count of selected lines per FILE")
	pflag.BoolVarP(&input.Flags.Flagi, "ignore-case", "i", false, "ignore case distinctions in patterns and data")
	pflag.BoolVarP(&input.Flags.Flagv, "invert-match", "v", false, "select non-matching lines")
	pflag.BoolVarP(&input.Flags.FlagF, "fixed-strings", "F", false, "PATTERNS are strings")
	pflag.BoolVarP(&input.Flags.Flagn, "line-number", "n", false, "print line number with output lines")

	pflag.Parse()

	for _, contextLines := range [3]int{input.Flags.FlagA, input.Flags.FlagB, input.Flags.FlagC} {
		if contextLines < 0 {
			input.CurrLineNumber = contextLines // store the invalid context value to report in logFatal
			return input, ErrInvalidContext
		}
	}

	args := pflag.Args()
	if len(args) == 0 {
		return input, ErrNoArgs
	}
	pattern := args[0]

	if input.Flags.FlagF {
		pattern = regexp.QuoteMeta(pattern)
	}

	if input.Flags.Flagi {
		pattern = "(?i)" + pattern
	}

	input.Pattern = regexp.MustCompile(pattern)

	if input.Flags.FlagC > 0 {
		input.Flags.FlagA, input.Flags.FlagB = input.Flags.FlagC, input.Flags.FlagC
	}

	if input.Flags.FlagA > 0 {
		input.LinesAfter = input.Flags.FlagA
	}

	if input.Flags.FlagB > 0 {
		input.LinesBefore = buffer.NewBuffer(input.Flags.FlagB)
	}

	input.FileNames = args[1:]

	return input, nil

}
