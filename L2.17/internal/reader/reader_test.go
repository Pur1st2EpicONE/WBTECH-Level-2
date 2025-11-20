package reader_test

import (
	"context"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"L2.17/internal/reader"
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

func TestRead_SimpleData(t *testing.T) {

	ln, addr := startEchoServer(t)
	defer func() { _ = ln.Close() }()

	conn, err := net.Dial("tcp", addr)
	assert.NoError(t, err)
	defer func() { _ = conn.Close() }()

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	var exitCode atomic.Int32

	wg.Add(1)
	go reader.Read(ctx, cancel, conn, &wg, &exitCode)

	msg := "qweqweqwe\n"
	_, err = conn.Write([]byte(msg))
	assert.NoError(t, err)

	time.Sleep(50 * time.Millisecond)
	cancel()
	wg.Wait()

	assert.Equal(t, int32(0), exitCode.Load())

}

func TestRead_Cancel(t *testing.T) {

	ln, addr := startEchoServer(t)
	defer func() { _ = ln.Close() }()

	conn, err := net.Dial("tcp", addr)
	assert.NoError(t, err)
	defer func() { _ = conn.Close() }()

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	var exitCode atomic.Int32

	wg.Add(1)
	go reader.Read(ctx, cancel, conn, &wg, &exitCode)

	time.Sleep(50 * time.Millisecond)
	cancel()
	wg.Wait()

	assert.Equal(t, int32(0), exitCode.Load())

}

func TestRead_EOF(t *testing.T) {

	conn1, conn2 := net.Pipe()
	defer func() { _ = conn2.Close() }()
	_ = conn1.Close()

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	var exitCode atomic.Int32

	wg.Add(1)
	go reader.Read(ctx, cancel, conn2, &wg, &exitCode)

	wg.Wait()
	assert.Equal(t, int32(1), exitCode.Load())

}

func TestRead_ReadError(t *testing.T) {

	conn := new(mockConn)
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	var exitCode atomic.Int32

	wg.Add(1)
	go reader.Read(ctx, cancel, conn, &wg, &exitCode)

	wg.Wait()
	assert.Equal(t, int32(1), exitCode.Load())

}

type mockConn struct{}

func (b *mockConn) Read(p []byte) (int, error)         { return 0, io.ErrUnexpectedEOF }
func (b *mockConn) Write(p []byte) (int, error)        { return len(p), nil }
func (b *mockConn) Close() error                       { return nil }
func (b *mockConn) LocalAddr() net.Addr                { return nil }
func (b *mockConn) RemoteAddr() net.Addr               { return nil }
func (b *mockConn) SetDeadline(t time.Time) error      { return nil }
func (b *mockConn) SetReadDeadline(t time.Time) error  { return nil }
func (b *mockConn) SetWriteDeadline(t time.Time) error { return nil }
