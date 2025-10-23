// Package input implements argument parsing and validation for the cut utility.
// It processes flags, field selections, ranges, and input file names.
package input

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

// Errors from ParseArgs, checked in logFatal if returned.
var (
	ErrInvalidFieldValue  = errors.New("invalid field value")
	ErrInvalidFieldRange  = errors.New("invalid field range")
	ErrNoFieldsToShow     = errors.New("you must specify a list of fields")
	ErrDecreasingRange    = errors.New("invalid decreasing range")
	ErrNotNumberedFromOne = errors.New("fields are numbered from 1")
)

// Flags holds command-line flag values.
type Flags struct {
	Flagf string // Selected fields as string
	Flagd string // Delimiter
	Flags bool   // Only print lines containing delimiters
}

// FieldRange represents a range of fields.
type FieldRange struct {
	Start int // Range start
	End   int // Range end
}

// Data holds parsed input configuration.
type Data struct {
	Flags       Flags        // Parsed command-line flags
	Fields      []int        // List of field numbers
	Ranges      []FieldRange // List of ranges
	FileNames   []string     // List of input file names
	CurrentFile string       // Name of the currently processed file
}

// ParseArgs parses command-line arguments into a Data structure.
func ParseArgs() (*Data, error) {

	input := new(Data)

	pflag.StringVarP(&input.Flags.Flagf, "fields", "f", "", "select only these fields;\nalso print any line that contains no delimiter character, unless\nthe -s option is specified")
	pflag.StringVarP(&input.Flags.Flagd, "delimiter", "d", "\t", "use DELIM instead of TAB for field delimiter")
	pflag.BoolVarP(&input.Flags.Flags, "only-delimited", "s", false, "do not print lines not containing delimiters")

	pflag.Parse()

	if input.Flags.Flagf == "" {
		return input, ErrNoFieldsToShow
	}

	if err := parseFields(input); err != nil {
		return input, err // intentionally not wrapped: the error is handled by logFatal, which formats it in the style of GNU cut using errors.Is checks.
	}

	args := pflag.Args()
	input.FileNames = args

	return input, nil

}

// parseFields parses and validates field selections.
func parseFields(input *Data) error {

	fieldsMap := make(map[int]struct{})

	for field := range strings.SplitSeq(input.Flags.Flagf, ",") {

		if field == "" {
			return ErrNotNumberedFromOne
		}

		if strings.Contains(field, "--") {
			return ErrInvalidFieldRange
		}

		if !strings.ContainsRune(field, '-') {
			if err := saveField(field, input, fieldsMap); err != nil {
				return err
			}
			continue
		}

		if err := saveRange(field, input, fieldsMap); err != nil {
			return err
		}

	}

	sortFields(&input.Fields, fieldsMap)
	return nil

}

// saveField stores a single field number.
func saveField(field string, input *Data, fieldsMap map[int]struct{}) error {

	filedNumber, err := strconv.Atoi(field)
	if err != nil {
		input.Flags.Flagf = field // store the invalid field value to report in logFatal
		return ErrInvalidFieldValue
	}

	if filedNumber == 0 {
		return ErrNotNumberedFromOne
	}

	fieldsMap[filedNumber] = struct{}{}
	return nil

}

// saveRange stores a field range.
func saveRange(field string, input *Data, fieldsMap map[int]struct{}) error {

	fields := strings.SplitN(field, "-", 2)

	rangeStart, err := getRangeStart(fields, input)
	if err != nil {
		return err
	}

	rangeEnd, err := getRangeEnd(fields, rangeStart, input)
	if err != nil {
		return err
	}

	for field := rangeStart; field <= rangeEnd; field++ {
		fieldsMap[field] = struct{}{}
	}

	return nil

}

// getRangeStart parses the start of a field range.
func getRangeStart(fields []string, input *Data) (int, error) {

	if fields[0] != "" { // -N

		rangeStart, err := strconv.Atoi(fields[0])
		if err != nil {
			input.Flags.Flagf = strings.Join(fields, "")
			return 0, ErrInvalidFieldValue
		}

		if rangeStart == 0 {
			input.Flags.Flagf = fields[0]
			return 0, ErrNotNumberedFromOne
		}

		if rangeStart < 0 {
			input.Flags.Flagf = fields[0]
			return 0, ErrInvalidFieldRange
		}

		return rangeStart, nil

	}

	return 1, nil // fields are numbered from 1 by default

}

// getRangeEnd parses the end of a field range.
func getRangeEnd(fields []string, rangeStart int, input *Data) (int, error) {

	if fields[1] == "" { // N-
		input.Ranges = append(input.Ranges, FieldRange{Start: rangeStart, End: -1})
		return -1, nil
	}

	rangeEnd, err := strconv.Atoi(fields[1])
	if err != nil || rangeEnd < 0 {
		return 0, ErrInvalidFieldRange
	}

	if rangeStart > rangeEnd {
		return 0, ErrDecreasingRange
	}

	return rangeEnd, nil

}

// sortFields sorts and stores unique fields.
func sortFields(fields *[]int, fieldsMap map[int]struct{}) {

	for fieled := range fieldsMap {
		if fieled >= 1 {
			*fields = append(*fields, fieled)
		}
	}

	sort.Ints(*fields)

}
