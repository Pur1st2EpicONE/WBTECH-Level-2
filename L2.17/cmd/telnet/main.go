// Package main provides the entry point for the telnet client application.
// It delegates all connection logic and runtime orchestration to the telnet package.
package main

import (
	"L2.17/internal/telnet"
)

func main() {

	telnet.Telnet()

}
