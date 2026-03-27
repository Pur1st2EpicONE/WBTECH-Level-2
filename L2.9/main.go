package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrNoCharacters = errors.New("failed to unpack string — no letters or characters found")
var ErrDigitWithoutChar = errors.New("failed to unpack string — digit without preceding character")

var strs = []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`}

func main() {

	for _, str := range strs {
		fmt.Println(unpack(str))
	}

}

// unpack expands a string according to a simple repetition rule.
// It returns the unpacked string or an error if the format is invalid.
func unpack(str string) (string, error) {

	if str == "" {
		return "", nil
	}

	if _, err := strconv.Atoi(str); err == nil {
		return "", ErrNoCharacters
	}

	var res strings.Builder
	runes := []rune(str)
	i := 0

	for i < len(runes) {

		var ch rune
		if runes[i] == '\\' {
			i++
			if i >= len(runes) {
				res.WriteRune('\\')
				break
			}
			ch = runes[i]
		} else if unicode.IsDigit(runes[i]) {
			return "", ErrDigitWithoutChar
		} else {
			ch = runes[i]
		}
		i++

		count, newI := getRepeatCount(runes, i)
		i = newI

		for range count {
			res.WriteRune(ch)
		}

	}

	return res.String(), nil

}

// getRepeatCount reads a decimal number starting at position start in runes
// and returns the integer value and the next index. If no digit is found,
// it returns (1, start) – meaning a default repeat count of 1.
func getRepeatCount(runes []rune, start int) (int, int) {

	if start >= len(runes) || !unicode.IsDigit(runes[start]) {
		return 1, start
	}

	var num strings.Builder
	i := start
	for i < len(runes) && unicode.IsDigit(runes[i]) {
		num.WriteRune(runes[i])
		i++
	}

	count, err := strconv.Atoi(num.String())
	if err != nil {
		panic(err) // theoretically impossible
	}

	return count, i

}
