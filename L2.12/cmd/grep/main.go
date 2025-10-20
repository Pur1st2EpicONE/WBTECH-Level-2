// Package main provides the entry point for the custom grep utility.
// It delegates all flag parsing, file processing, and matching logic
// to the internal grep package, keeping the main function minimal.
package main

import "L2.12/internal/grep"

func main() {

	grep.Grep()

}
