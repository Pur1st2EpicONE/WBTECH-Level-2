package main

import (
	"testing"
	"time"
)

func sig(after time.Duration) <-chan any {
	c := make(chan any)
	go func() {
		time.Sleep(after)
		close(c)
	}()
	time.Sleep(10 * time.Millisecond)
	return c
}

func assertDuration(t *testing.T, got time.Duration, want time.Duration, tolerance time.Duration) {
	if got < want-tolerance || got > want+tolerance {
		t.Fatalf("expected around %v, got %v", want, got)
	}
}

func TestOr1(t *testing.T) {

	start := time.Now()
	<-or1(
		sig(2*time.Second),
		sig(100*time.Millisecond),
		sig(1*time.Second),
	)

	elapsed := time.Since(start)
	assertDuration(t, elapsed, 100*time.Millisecond, 50*time.Millisecond)

}

func TestOr2(t *testing.T) {

	start := time.Now()

	<-or2(
		sig(500*time.Millisecond),
		sig(2*time.Second),
	)

	elapsed := time.Since(start)
	assertDuration(t, elapsed, 500*time.Millisecond, 50*time.Millisecond)

}

func TestOr3(t *testing.T) {

	start := time.Now()

	<-or3(
		sig(3*time.Second),
		sig(200*time.Millisecond),
		sig(1*time.Second),
	)

	elapsed := time.Since(start)
	assertDuration(t, elapsed, 200*time.Millisecond, 50*time.Millisecond)

}

func TestSingleChannel(t *testing.T) {

	c := sig(300 * time.Millisecond)

	testOr := func(name string, orFunc func(...<-chan any) <-chan any) {
		start := time.Now()
		<-orFunc(c)
		elapsed := time.Since(start)
		t.Logf("%s elapsed: %v", name, elapsed)
		assertDuration(t, elapsed, 300*time.Millisecond, 50*time.Millisecond)
	}

	testOr("or1", or1)

	c = sig(300 * time.Millisecond)
	testOr("or2", or2)

	c = sig(300 * time.Millisecond)
	testOr("or3", or3)

}

func TestNoChannels(t *testing.T) {

	if ch := or1(); ch != nil {
		t.Fatalf("expected nil for or1 with no channels")
	}

	if ch := or2(); ch != nil {
		t.Fatalf("expected nil for or2 with no channels")
	}

	if ch := or3(); ch != nil {
		t.Fatalf("expected nil for or3 with no channels")
	}

}
