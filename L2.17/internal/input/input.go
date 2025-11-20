// Package input provides parsing and validation of command-line arguments
// for configuring and establishing a telnet connection.
package input

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/spf13/pflag"
)

// These error messages intentionally begin with uppercase letters.
// Staticcheck warns about this naming style, but the formatting is
// preserved to reflect the conventions of the original telnet utility,
// and the warnings are suppressed using nolint directives.
var (
	//nolint:staticcheck
	ErrHostNameLookUp = errors.New("Host name lookup failure")
	//nolint:staticcheck
	ErrInvalidPort = errors.New("Invalid port name")
	//nolint:staticcheck
	ErrWrongUsage = errors.New("Usage: telnet [--timeout DURATION] [host-name [port]]")
)

// Data holds parsed command-line input required to establish a telnet connection.
type Data struct {
	Host    string        // Host is the hostname provided by the user (domain name or IP literal).
	IP      string        // IP is the resolved IP address for Host.
	Port    string        // Port is the TCP port as a string.
	Timeout time.Duration // Timeout is the connection timeout duration.
}

// ParseArgs parses command-line flags and positional arguments, returning
// a Data struct with the resolved host, IP, port, and timeout, or an error
// if parsing or validation fails.
func ParseArgs() (*Data, error) {

	input := new(Data)

	pflag.DurationVar(&input.Timeout, "timeout", 10*time.Second, "connection timeout")
	pflag.Parse()

	if err := validate(input, pflag.Args()); err != nil {
		return nil, err
	}

	return input, nil

}

// validate validates positional arguments and fills the provided Data with host, IP and port.
// It returns an error when arguments are missing or invalid.
func validate(input *Data, args []string) error {

	if len(args) == 0 {
		return ErrWrongUsage
	} else if len(args) == 1 {
		input.Port = "23" // default telnet behavior, uses port 23 if none is specified
	} else {
		port, err := strconv.Atoi(args[1])
		if err != nil || port < 1 || port > 65535 {
			return fmt.Errorf("%v '%s'", ErrInvalidPort, args[1])
		}
		input.Port = args[1]
	}

	address, err := net.LookupIP(args[0])
	if err != nil {
		return fmt.Errorf("%s: %v", args[0], ErrHostNameLookUp)
	}

	input.IP = address[0].String()
	input.Host = args[0]

	return nil

}
