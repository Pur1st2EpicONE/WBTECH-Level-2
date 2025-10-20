// Package grep implements a simplified version of the GNU grep utility.
// It supports common flags, colored output, context lines, and multiple input files.
package grep

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"L2.12/internal/input"
)

// Exit codes and ANSI color constants for grep output formatting.
const (
	matchNotFound = 1
	internalError = 2

	red     = "\033[31m"
	green   = "\033[32m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	bold    = "\033[1m"
	reset   = "\033[0m"
)

// Grep executes the main grep logic: parses args, processes files or stdin.
func Grep() {

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

	if !data.MatchFound {
		os.Exit(matchNotFound)
	}

}

// processFiles iterates over files and processes each one.
func processFiles(data *input.Data) error {

	for _, fileName := range data.FileNames {
		data.CurrentFile = fileName
		if err := processFile(data); err != nil {
			return err // intentionally not wrapped: the error is handled by logFatal, which formats it in the style of GNU grep using errors.Is checks.
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

	if data.Flags.Flagc {
		fmt.Println(data.CurrLineNumber)
	}

}

// processFile opens a file, scans it line by line, and applies grep logic.
func processFile(data *input.Data) error {

	file, err := os.Open(data.CurrentFile)
	if err != nil {
		return err // intentionally not wrapped: the error is handled by logFatal, which formats it in the style of GNU grep using errors.Is checks.
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processLine(scanner.Text(), data)
	}

	if data.Flags.Flagc {
		if len(data.FileNames) > 1 {
			fmt.Printf("%s%s%s:%s%d\n", magenta, data.CurrentFile, cyan, reset, data.CurrLineNumber)
		} else {
			fmt.Println(data.CurrLineNumber)
		}

	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}

	if err := scanner.Err(); err != nil {
		return err // intentionally not wrapped: the error is handled by logFatal, which formats it in the style of GNU grep using errors.Is checks.
	}

	data.CurrLineNumber = 0
	return nil

}

// processLine processes a single line of input based on grep flags.
func processLine(line string, data *input.Data) {

	if !data.Flags.Flagc {
		data.CurrLineNumber++
	}

	switch {
	case data.Pattern.MatchString(line):
		handleMatch(line, data)
	case data.Flags.Flagv:
		handleInvertedMatch(line, data.CurrLineNumber, data)
	case data.LinesAfter > 0 && data.MatchFound:
		handleAfterContext(line, data.CurrLineNumber, data)
	default:
		handleBeforeContext(line, data.CurrLineNumber, data)
	}

}

// handleMatch processes a matching line, including context and highlighting.
func handleMatch(line string, data *input.Data) {

	data.MatchFound = true

	if data.Flags.Flagc && !data.Flags.Flagv {
		data.CurrLineNumber++
		return
	}

	if needsSeparator(data) {
		fmt.Printf("%s--%s\n", cyan, reset)
	}

	if data.Flags.FlagB > 0 && data.LinesBefore.Size > 0 {
		printBufferedLines(data)
	}

	if data.Flags.FlagA > 0 {
		data.LinesAfter = data.Flags.FlagA
	}

	line = string(data.Pattern.ReplaceAllFunc([]byte(line), highlightMatch))

	if !data.Flags.Flagv {
		printLine(true, data, line, data.CurrLineNumber, ":")
	}

	data.LastPrintedLineNumber = data.CurrLineNumber

}

// needsSeparator checks if a separator line should be printed.
func needsSeparator(data *input.Data) bool {

	if data.LastPrintedLineNumber == 0 || data.Flags.FlagA == 0 && data.Flags.FlagB == 0 {
		return false
	}
	return data.CurrLineNumber-data.LastPrintedLineNumber > data.Flags.FlagB+1

}

// printBufferedLines prints stored lines before the current match.
func printBufferedLines(data *input.Data) {

	for offset, index := 0, 0; offset < data.LinesBefore.Size; offset++ {

		index = (data.LinesBefore.Head + offset) % len(data.LinesBefore.Buffer) // common way to compute the actual element index in a circular queue

		if data.LinesBefore.Printed[index] {
			continue
		}

		printLine(false, data, data.LinesBefore.Buffer[index], data.LinesBefore.Numbers[index], "-")
		data.LinesBefore.Printed[index] = true

	}

}

// printLine prints a line with proper prefixes, colors, and formatting.
func printLine(match bool, data *input.Data, line string, num int, sep string) {

	if data.Flags.Flagn {

		if len(data.FileNames) > 1 {
			if match {
				fmt.Printf("%s%s%s:%s%d%s%s%s%s\n", magenta, data.CurrentFile, cyan, green, num, cyan, sep, reset, line)
			} else {
				fmt.Printf("%s%s%s-%s%d%s%s%s%s\n", magenta, data.CurrentFile, cyan, green, num, cyan, sep, reset, line)
			}
		} else {
			fmt.Printf("%s%d%s%s%s%s\n", green, num, cyan, sep, reset, line)
		}

	} else {

		if len(data.FileNames) > 1 {
			fmt.Printf("%s%s%s:%s%s\n", magenta, data.CurrentFile, cyan, reset, line)
		} else {
			fmt.Println(line)
		}

	}

}

// highlightMatch adds ANSI highlighting to matched text.
func highlightMatch(b []byte) []byte {
	return []byte(bold + red + string(b) + reset)
}

// handleInvertedMatch processes lines that do not match when -v is used.
func handleInvertedMatch(line string, current int, data *input.Data) {

	if data.Flags.Flagc {
		data.CurrLineNumber++
		return
	}

	printLine(false, data, line, current, ":")
	data.LastPrintedLineNumber = current

	if data.Flags.FlagB > 0 && data.LinesBefore != nil {
		_ = data.LinesBefore.Enqueue(line, current)
	}

}

// handleAfterContext prints lines following a match for -A.
func handleAfterContext(line string, current int, data *input.Data) {

	if data.Flags.FlagB > 0 && data.LinesBefore != nil {
		index := data.LinesBefore.Enqueue(line, current)
		data.LinesBefore.Printed[index] = true
	}

	printLine(false, data, line, current, "-")
	data.LastPrintedLineNumber = current

	data.LinesAfter--

}

// handleBeforeContext enqueues lines before a match for -B.
func handleBeforeContext(line string, current int, data *input.Data) {
	if data.Flags.FlagB > 0 && data.LinesBefore != nil {
		_ = data.LinesBefore.Enqueue(line, current)
	}
}

// logFatal prints formatted error messages consistent with GNU grep behavior
func logFatal(err error, data *input.Data) {

	switch {
	case errors.Is(err, input.ErrInvalidContext):
		fmt.Fprintf(os.Stderr, "grep: %d: invalid context length argument\n", data.CurrLineNumber)
	case errors.Is(err, input.ErrNoArgs):
		fmt.Fprintf(os.Stderr, "Usage: grep [OPTION]... PATTERNS [FILE]...\n")
		fmt.Fprintf(os.Stderr, "Try 'grep --help' for more information.\n")
	case errors.Is(err, os.ErrNotExist):
		fmt.Fprintf(os.Stderr, "grep: %s: No such file or directory\n", data.CurrentFile)
	default:
		fmt.Fprintf(os.Stderr, "grep: fatal error: %v\n", err)
	}

	os.Exit(internalError) // Exit with code 2 for internal errors, matching GNU grep behavior

}
