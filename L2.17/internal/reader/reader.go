// Package reader contains functionality for reading data from a remote TCP connection
// and writing it to stdout until the context is cancelled or an error occurs.
package reader

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

const readBufferSize = 2048
const readTimeout = 100 * time.Millisecond

// Read reads data from the TCP connection and writes it to stdout until the context is cancelled
// or a fatal error occurs. It signals completion by calling wg.Done() and cancels the context on exit.
func Read(ctx context.Context, cancel context.CancelFunc, conn net.Conn, wg *sync.WaitGroup, exitCode *atomic.Int32) {

	defer wg.Done()
	defer cancel()

	reader := bufio.NewReader(conn)
	buffer := make([]byte, readBufferSize)

	for {

		if ctx.Err() != nil {
			return
		}

		if err := conn.SetReadDeadline(time.Now().Add(readTimeout)); err != nil {
			fmt.Fprintf(os.Stderr, "telnet: failed to set the deadline for future Read call: %v\n", err)
			exitCode.CompareAndSwap(0, 1)
			return
		}

		n, err := reader.Read(buffer)
		if err != nil {

			if errors.Is(err, io.EOF) {
				fmt.Fprintln(os.Stderr, "Connection closed by foreign host.")
				exitCode.CompareAndSwap(0, 1)
				return
			}

			var netErr net.Error
			if errors.As(err, &netErr) {
				if netErr.Timeout() {
					continue
				}
			}

			if !errors.Is(err, os.ErrDeadlineExceeded) && !errors.Is(err, net.ErrClosed) {
				fmt.Fprintf(os.Stderr, "telnet: read error from remote host %s: %v\n", conn.RemoteAddr(), err)
				exitCode.CompareAndSwap(0, 1)
				return
			}

		}

		if _, err := fmt.Fprint(os.Stdout, string(buffer[:n])); err != nil {
			fmt.Fprintf(os.Stderr, "failed to print to stdout: %v\n", err)
			exitCode.CompareAndSwap(0, 1)
			return
		}

	}

}
