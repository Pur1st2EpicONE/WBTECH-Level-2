// Package main is the entry point of the wget-like scraper application.
// It initializes and runs the wget logic from the internal/wget package.
package main

import (
	"L2.16/internal/wget"
)

// main runs the wget process.
func main() {

	wget.Wget()

}
