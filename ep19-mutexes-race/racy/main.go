// A program with a race condition, on purpose. Run it with: go run ./racy
// Then catch the bug with: go run -race ./racy
package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)
}
