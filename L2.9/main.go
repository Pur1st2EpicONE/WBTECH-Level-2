package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	fmt.Println(unpack("a4bc2d5"))
	fmt.Println(unpack(`qwe\4\5`))
	fmt.Println(unpack(`qwe\45\`))
	fmt.Println(unpack(`a2bc4\7d!2\\2`))

}

// unpack takes a string with run-length encoding and returns its unpacked form.
func unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	_, err := strconv.Atoi(str)
	if err == nil {
		return "", errors.New("failed to unpack string — no letters or characters found")
	}
	var res, num strings.Builder // res — result string, num — buffer for digits
	runes := []rune(str)
	i := 0
	for i < len(runes) {
		if runes[i] == '\\' {
			if i+1 >= len(runes) { // processing single backslash at the very end (treated as is)
				res.WriteRune('\\')
				i++
				continue
			}
			if runes[i+1] != '\\' { // processing single backslash followed by something other than other backslash
				if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) { // processing escaped digits
					repeats, err := toInt(runes, &num, i+2)
					if err != nil {
						return str, fmt.Errorf("failed to parse number of letter repeats: %v", err)
					}
					for range repeats {
						res.WriteRune(runes[i+1])
					}
					i += 2 + num.Len()
				} else {
					res.WriteRune(runes[i+1]) // processing escaped characters (they are treated as is)
					i += 2
				}
				continue
			}
			repeats, err := toInt(runes, &num, i+2) // processing double backslash that is possibly followed by a repeat count
			if err != nil {
				return str, fmt.Errorf("failed to parse number of letter repeats: %v", err)
			}
			if num.Len() == 0 { // processing "\\" (just writing one backslash to result string)
				res.WriteRune('\\')
				i += 2
			} else {
				for range repeats { // processing backslashes with mulitplication (\\3 —> \\\)
					res.WriteRune('\\')
				}
				i += 2 + num.Len()
			}
			continue
		}
		if !unicode.IsDigit(runes[i]) { // processing regular non-digit characters
			if i == len(runes)-1 || !unicode.IsDigit(runes[i+1]) || runes[i+1] == '\\' { // if next rune is not a digit just writing it as-is
				res.WriteRune(runes[i])
			} else {
				repeats, err := toInt(runes, &num, i+1) // if next rune is a digit, building the number and repeating the rune N times
				if err != nil {
					return str, fmt.Errorf("failed to parse number of letter repeats: %v", err)
				}
				for range repeats {
					res.WriteRune(runes[i])
				}
				i += num.Len()
			}
		}
		i++
	}
	return res.String(), nil
}

// toInt extracts a number from the rune slice starting at i (position after the backslash).
// If no number is found, returns 1 (default repeat count).
func toInt(runes []rune, num *strings.Builder, i int) (int, error) {
	num.Reset()
	for i < len(runes) && unicode.IsDigit(runes[i]) { // processing consecutive digits
		num.WriteRune(runes[i])
		i++
	}
	if num.Len() == 0 { // if no digits found, repeat the character once
		return 1, nil
	}
	repeats, err := strconv.Atoi(num.String())
	if err != nil {
		return 0, fmt.Errorf("atoi failed to convert string to int: %w", err)
	}
	return repeats, nil
}
