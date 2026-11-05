package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// run starts 1,000 goroutines that each call add 1,000 times.
func run(add func()) {
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				add()
			}
		}()
	}
	wg.Wait()
}

// Cache allows many readers at once, or one writer.
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func main() {
	// sync.Mutex: only one goroutine at a time gets past Lock.
	count := 0
	var mu sync.Mutex
	run(func() {
		mu.Lock()
		count++
		mu.Unlock()
	})
	fmt.Println("mutex:", count)

	// sync/atomic: for simple counters.
	var total atomic.Int64
	run(func() {
		total.Add(1)
	})
	fmt.Println("atomic:", total.Load())

	// sync.RWMutex
	cache := Cache{data: map[string]string{}}
	cache.Set("lang", "Go")
	fmt.Println(cache.Get("lang"))
}
