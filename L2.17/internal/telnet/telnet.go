// Package telnet implements the high-level orchestration of the telnet client:
// argument parsing, TCP connection establishment, goroutine management,
// signal handling, and graceful shutdown.
package telnet

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"

	"L2.17/internal/input"
	"L2.17/internal/reader"
	"L2.17/internal/writer"
)

// Telnet runs the telnet client: it parses arguments, connects to the remote host,
// starts reader/writer goroutines and an interrupt handler, waits for completion,
// and exits with the appropriate exit code.
func Telnet() {

	data, err := input.ParseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Trying %s...\n", data.IP)

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(data.Host, data.Port), data.Timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "telnet: Unable to connect to remote host: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Connected to %s.\nEscape character is '^D'.\n", data.Host)

	var wg sync.WaitGroup
	var exitCode atomic.Int32
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(3)

	go reader.Read(ctx, cancel, conn, &wg, &exitCode)
	go writer.Write(ctx, cancel, conn, &wg, &exitCode)
	go interruptHandler(ctx, sigCh, cancel, &wg, &exitCode)

	wg.Wait()

	if err := conn.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "telnet: failed to close TCP connection: %v\n", err)
		exitCode.CompareAndSwap(0, 1)
	}

	os.Exit(int(exitCode.Load()))

}

// interruptHandler listens for OS interrupt/terminate signals and cancels the provided context.
// In real telnet, Ctrl+C does not immediately terminate the client. Here, although a graceful
// shutdown is performed, exiting due to SIGINT/SIGTERM is not normal, so the exit code is set to 1.
func interruptHandler(ctx context.Context, sigCh chan os.Signal, cancel context.CancelFunc, wg *sync.WaitGroup, exitCode *atomic.Int32) {

	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	case <-sigCh:
		fmt.Println()
		exitCode.CompareAndSwap(0, 1)
		cancel()
	}

}
