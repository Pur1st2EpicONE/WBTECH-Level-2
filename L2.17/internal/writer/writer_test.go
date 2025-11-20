package writer_test

import (
	"context"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"L2.17/internal/writer"
)

func startEchoServer(t *testing.T) (net.Listener, string) {

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatalf("failed to start server: %v", err)
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				defer func() { _ = c.Close() }()
				buf := make([]byte, 2048)
				for {
					n, err := c.Read(buf)
					if err != nil {
						return
					}
					_, _ = c.Write(buf[:n])
				}
			}(conn)
		}
	}()

	return ln, ln.Addr().String()

}

func TestWrite_Cancel(t *testing.T) {

	ln, addr := startEchoServer(t)
	defer func() { _ = ln.Close() }()

	conn, err := net.Dial("tcp", addr)
	assert.NoError(t, err)
	defer func() { _ = conn.Close() }()

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	var exitCode atomic.Int32

	wg.Add(1)
	go writer.Write(ctx, cancel, conn, &wg, &exitCode)

	time.Sleep(50 * time.Millisecond)
	cancel()
	wg.Wait()

	assert.Equal(t, int32(0), exitCode.Load())

}

func TestWrite_SendAndReceive(t *testing.T) {

	ln, addr := startEchoServer(t)
	defer func() { _ = ln.Close() }()

	conn, err := net.Dial("tcp", addr)
	assert.NoError(t, err)
	defer func() { _ = conn.Close() }()

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	var exitCode atomic.Int32

	wg.Add(1)
	go writer.Write(ctx, cancel, conn, &wg, &exitCode)

	go func() {
		writer.StdInReader(make(chan string), make(chan error))
	}()

	msg := "hello\n"
	_, err = conn.Write([]byte(msg))
	assert.NoError(t, err)

	buf := make([]byte, len(msg))
	n, err := conn.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, msg, string(buf[:n]))

	cancel()
	wg.Wait()
	assert.Equal(t, int32(0), exitCode.Load())

}

func TestStdInReader_EOF(t *testing.T) {

	inputCh := make(chan string, 1)
	errCh := make(chan error, 1)

	_, w := net.Pipe()

	go func() {
		defer func() { _ = w.Close() }()
		writer.StdInReader(inputCh, errCh)
	}()

	_ = w.Close()

	select {
	case err := <-errCh:
		assert.ErrorIs(t, err, io.EOF)
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout waiting for EOF")
	}

}
