// Package cut implements a simplified version of the Unix cut utility.
// It processes input lines, applies field selection, and prints results.
package cut

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"L2.13/internal/input"
)

const internalError = 1

// Cut executes the main cut logic: parses args, processes files or stdin.
func Cut() {

	data, err := input.ParseArgs()
	if err != nil {
		logFatal(err, data)
	}

	if len(data.FileNames) > 0 {
		if err := processFiles(data); err != nil {
			logFatal(err, data)
		}
	} else {
		processStdIn(data)
	}

}

// processFiles iterates over files and processes each one.
func processFiles(data *input.Data) error {

	for _, fileName := range data.FileNames {
		data.CurrentFile = fileName
		if err := processFile(data); err != nil {
			return err // intentionally not wrapped: the error is handled by logFatal, which formats it in the style of GNU cut using errors.Is checks.
		}
	}

	return nil

}

// processStdIn handles input from stdin when no files are given.
func processStdIn(data *input.Data) {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		processLine(scanner.Text(), data)
	}

}

// processFile opens a file, scans it line by line, and applies cut logic.
func processFile(data *input.Data) error {

	file, err := os.Open(data.CurrentFile)
	if err != nil {
		return err // intentionally not wrapped: the error is handled by logFatal, which formats it in the style of GNU cut using errors.Is checks.
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processLine(scanner.Text(), data)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}

	if err := scanner.Err(); err != nil {
		return err // intentionally not wrapped: the error is handled by logFatal, which formats it in the style of GNU cut using errors.Is checks.
	}

	return nil

}

// processLine processes a single line of input based on cut flags.
func processLine(line string, data *input.Data) {

	allFields := strings.Split(line, data.Flags.Flagd)

	if !strings.Contains(line, data.Flags.Flagd) {
		if !data.Flags.Flags {
			fmt.Println(line)
		}
		return
	}

	fieldsToPrint := make([]string, 0, len(data.Fields))

	for _, fieldNumber := range data.Fields {
		if fieldNumber <= len(allFields) {
			fieldsToPrint = append(fieldsToPrint, allFields[fieldNumber-1])
		}
	}

	for _, currentRange := range data.Ranges {
		if currentRange.End == -1 {
			fieldsToPrint = append(fieldsToPrint, allFields[currentRange.Start-1:]...) // add all fields starting from the range start to the end
		}
	}

	fmt.Println(strings.Join(fieldsToPrint, data.Flags.Flagd))

}

// logFatal prints formatted error messages consistent with GNU cut behavior.
func logFatal(err error, data *input.Data) {

	switch {
	case errors.Is(err, input.ErrNoFieldsToShow):
		fmt.Fprintf(os.Stderr, "cut: you must specify a list of bytes, characters, or fields\n")
	case errors.Is(err, input.ErrNotNumberedFromOne):
		fmt.Fprintf(os.Stderr, "cut: fields are numbered from 1\n")
	case errors.Is(err, input.ErrInvalidFieldValue):
		fmt.Fprintf(os.Stderr, "cut: invalid field value ‘%s’\n", data.Flags.Flagf)
	case errors.Is(err, input.ErrDecreasingRange):
		fmt.Fprintf(os.Stderr, "cut: invalid decreasing range\n")
	case errors.Is(err, input.ErrInvalidFieldRange):
		fmt.Fprintf(os.Stderr, "cut: invalid field range\n")
	default:
		if errors.Is(err, os.ErrNotExist) {
			fmt.Fprintf(os.Stderr, "cut: %s: No such file or directory\n", data.CurrentFile)
		} else {
			fmt.Fprintf(os.Stderr, "cut: fatal error: %v\n", err)
		}
		os.Exit(internalError)
	}

	fmt.Fprintf(os.Stderr, "Try 'cut --help' for more information.\n")
	os.Exit(internalError)

}
