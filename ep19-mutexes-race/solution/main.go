// Solution to the episode 19 exercise. Prove it's safe with: go run -race ./solution
package main

import (
	"fmt"
	"sync"
)

// Counter is safe to use from many goroutines at once.
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc adds one.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current count.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter.Value())
}
