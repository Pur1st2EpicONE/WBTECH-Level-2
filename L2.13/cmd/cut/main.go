// Package main provides the entry point for the custom cut utility.
// It delegates all flag parsing, field selection, and input processing
// to the internal cut package, keeping the main function minimal.
package main

import "L2.13/internal/cut"

func main() {

	cut.Cut()

}
