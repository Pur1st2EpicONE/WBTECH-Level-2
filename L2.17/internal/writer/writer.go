// Package writer provides functionality for reading user input from stdin,
// sending it to a remote TCP connection, and handling related errors
// until the context is cancelled or EOF is received.
package writer

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

const stdinBufferSize = 2048

// Write consumes strings from inputCh and sends them to the remote connection.
// It handles stdin read errors coming from errCh, flushes the writer, and sets exitCode on failure.
// The function returns after calling wg.Done() and cancelling the context via defer.
func Write(ctx context.Context, cancel context.CancelFunc, conn net.Conn, wg *sync.WaitGroup, exitCode *atomic.Int32) {

	defer wg.Done()
	defer cancel()

	errCh := make(chan error)
	inputCh := make(chan string)

	go StdInReader(inputCh, errCh)

	writer := bufio.NewWriter(conn)

	for {

		select {

		case <-ctx.Done():
			return
		case err := <-errCh:
			if errors.Is(err, io.EOF) {
				fmt.Println("^D\nConnection closed.")
			} else {
				fmt.Fprintf(os.Stderr, "telnet: read error from stdin: %v\n", err)
				exitCode.CompareAndSwap(0, 1)
			}
			return
		case line := <-inputCh:
			_, err := writer.WriteString(strings.ReplaceAll(line, "\n", "\r\n")) // some telnet servers expect CRLF (\r\n) endings
			if err != nil {
				fmt.Fprintf(os.Stderr, "telnet: failed to write to remote host: %v\n", err)
				exitCode.CompareAndSwap(0, 1)
				return
			}
			if err := writer.Flush(); err != nil {
				fmt.Fprintf(os.Stderr, "telnet: failed to flush data to remote host: %v\n", err)
				exitCode.CompareAndSwap(0, 1)
				return
			}

		}

	}

}

// StdInReader reads bytes from os.Stdin and sends them as strings to inputCh.
// Any read errors, including EOF, are forwarded to errCh. This goroutine does not
// listen to context cancellation because it blocks on stdin reads; this ensures
// that user input can be processed until EOF or an error occurs. By the time it
// exits, the TCP connection is already closed and other goroutines have finished,
// so there is no risk of goroutine leaks despite the blocking reads.
func StdInReader(inputCh chan string, errCh chan error) {

	buffer := make([]byte, stdinBufferSize)

	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			errCh <- err
			return
		} else if n > 0 {
			inputCh <- string(buffer[:n])
		}
	}

}
