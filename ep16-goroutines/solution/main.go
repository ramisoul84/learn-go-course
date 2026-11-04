// Solution to the episode 16 exercise. Run it with: go run ./solution
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://go.dev",
		"https://example.com",
		"https://github.com",
		"https://www.rust-lang.org",
		"https://pkg.go.dev",
	}

	start := time.Now()
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(url, "error:", err)
				return
			}
			defer resp.Body.Close()
			fmt.Println(url, resp.Status)
		}()
	}
	wg.Wait()
	fmt.Println("all done in", time.Since(start).Round(time.Millisecond))
}
