package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		channel := make(chan interface{})
		go func() {
			defer close(channel)
			time.Sleep(after)
		}()
		return channel
	}

	start := time.Now()
	<-or1(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("or1 done after %v\n", time.Since(start))

	start = time.Now()
	<-or2(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(2*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("or2 done after %v\n", time.Since(start))

	start = time.Now()
	<-or3(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(3*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("or3 done after %v\n", time.Since(start))
}

// or1 merges multiple done channels using a fan-out pattern.
func or1(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan any)
	go func() {
		var once sync.Once
		for _, channel := range channels {
			go func(channel <-chan any) {
				select {
				case <-channel:
					once.Do(func() { close(orDone) })
				case <-orDone:
				}
			}(channel)
		}
	}()
	return orDone
}

// or2 merges done channels using reflection and reflect.Select.
func or2(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan any)
	go func() {
		defer close(orDone)
		var cases []reflect.SelectCase
		for _, channel := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(channel),
			})
		}
		reflect.Select(cases)
	}()
	return orDone
}

// or3 combines done channels recursively using a divide-and-conquer approach.
func or3(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan any)
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			m := len(channels) / 2
			select {
			case <-or3(channels[:m]...):
			case <-or3(channels[m:]...):
			}
		}
	}()
	return orDone
}

/*
Output:
or1 done after 1.000112076s
or2 done after 2.00076626s
or3 done after 3.001129782s
*/
