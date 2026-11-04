package main

import (
	"fmt"
	"time"
)

// fetch pretends to download something, then sends the result down the channel.
func fetch(name string, d time.Duration, out chan string) {
	time.Sleep(d)
	out <- name + " finished"
}

// generate only sends, so its parameter is a send-only channel.
func generate(out chan<- int) {
	for i := 1; i <= 3; i++ {
		out <- i
	}
	close(out)
}

// show only receives, so its parameter is a receive-only channel.
func show(in <-chan int) {
	for n := range in {
		fmt.Println("got", n)
	}
}

func main() {
	// The program from the start of the episode.
	results := make(chan string)
	go fetch("slow", 300*time.Millisecond, results)
	go fetch("fast", 100*time.Millisecond, results)
	go fetch("medium", 200*time.Millisecond, results)
	for range 3 {
		fmt.Println(<-results)
	}

	// The simplest channel: one send, one receive.
	ch := make(chan string)
	go func() {
		ch <- "hello"
	}()
	msg := <-ch
	fmt.Println(msg)

	// A buffered channel holds values without a receiver, up to its capacity.
	buf := make(chan int, 2)
	buf <- 1
	buf <- 2
	fmt.Println(len(buf), cap(buf))
	fmt.Println(<-buf, <-buf)

	// close and range, with directional channels.
	nums := make(chan int)
	go generate(nums)
	show(nums)

	// Receiving from a closed channel gives the zero value, and ok is false.
	v, ok := <-nums
	fmt.Println(v, ok)
}
