// Solution to the episode 18 exercise. Run it with: go run ./solution
package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

// fetch makes a request that gives up when ctx ends.
func fetch(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	return nil
}

func main() {
	// A local test server that takes 5 seconds to answer.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(5 * time.Second):
		case <-r.Context().Done():
		}
	}))
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	start := time.Now()
	err := fetch(ctx, server.URL)
	fmt.Println("error:", err)
	fmt.Println("gave up after", time.Since(start).Round(100*time.Millisecond))
}
