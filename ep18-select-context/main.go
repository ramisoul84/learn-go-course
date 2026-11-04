package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// ask pretends to call a server, and answers on the returned channel after a delay.
// The channel has room for one value, so the goroutine can finish even if nobody listens.
func ask(name string, delay time.Duration) <-chan string {
	answer := make(chan string, 1)
	go func() {
		time.Sleep(delay)
		answer <- name + " answered"
	}()
	return answer
}

// work takes two seconds, unless its context ends first.
func work(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	// The first answer wins.
	a := ask("server-a", 300*time.Millisecond)
	b := ask("server-b", 100*time.Millisecond)
	select {
	case msg := <-a:
		fmt.Println(msg)
	case msg := <-b:
		fmt.Println(msg)
	}

	// A timeout with time.After.
	slow := ask("slow server", 2*time.Second)
	select {
	case msg := <-slow:
		fmt.Println(msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timed out after 500ms")
	}

	// A non-blocking check with default.
	inbox := make(chan string)
	select {
	case msg := <-inbox:
		fmt.Println(msg)
	default:
		fmt.Println("no message yet")
	}

	// context.WithTimeout: the deadline travels with ctx.
	ctx, cancel := context.WithTimeout(
		context.Background(), 500*time.Millisecond,
	)
	defer cancel()
	err := work(ctx)
	fmt.Println(err)
	fmt.Println(errors.Is(err, context.DeadlineExceeded))

	// context.WithCancel: you decide when to stop.
	ctx2, cancel2 := context.WithCancel(context.Background())
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel2()
	}()
	fmt.Println(work(ctx2))
}
