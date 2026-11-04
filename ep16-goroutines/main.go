package main

import (
	"fmt"
	"sync"
	"time"
)

// work pretends to be a slow job, like a download.
func work(name string) {
	time.Sleep(time.Second)
	fmt.Println(name, "done")
}

// oneByOne runs each job after the last one finishes.
func oneByOne() {
	work("A")
	work("B")
	work("C")
}

// atTheSameTime starts a goroutine for each job, and waits for them all.
func atTheSameTime() {
	var wg sync.WaitGroup
	for _, name := range []string{"A", "B", "C"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(name)
		}()
	}
	wg.Wait()
}

// many starts a huge number of goroutines, to show how cheap they are.
func many(n int) {
	var wg sync.WaitGroup
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
		}()
	}
	wg.Wait()
}

func main() {
	start := time.Now()
	oneByOne()
	fmt.Println("one by one:", time.Since(start).Round(time.Millisecond))

	start = time.Now()
	atTheSameTime()
	fmt.Println("at the same time:", time.Since(start).Round(time.Millisecond))

	start = time.Now()
	many(100_000)
	fmt.Println("100,000 goroutines:", time.Since(start).Round(time.Millisecond))
}
