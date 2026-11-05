// A concurrent website checker: a small pool of workers checks a list of sites.
package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"sync"
	"time"
)

const (
	workers = 3
	timeout = 3 * time.Second
)

// Result is what a worker found out about one site.
type Result struct {
	Site   string
	Status string
	Took   time.Duration
}

// check requests one site, and gives up after the timeout.
func check(site string) Result {
	start := time.Now()
	ctx, cancel := context.WithTimeout(
		context.Background(), timeout,
	)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet, "https://"+site, nil,
	)
	if err != nil {
		return Result{site, "bad URL", time.Since(start)}
	}
	resp, err := http.DefaultClient.Do(req)
	took := time.Since(start).Round(time.Millisecond)
	if errors.Is(err, context.DeadlineExceeded) {
		return Result{site, "TIMEOUT", took}
	}
	if err != nil {
		return Result{site, "DOWN", took}
	}
	defer resp.Body.Close()
	return Result{site, resp.Status, took}
}

// worker takes sites from jobs until the channel closes.
func worker(jobs <-chan string, results chan<- Result) {
	for site := range jobs {
		results <- check(site)
	}
}

func main() {
	sites := []string{
		"go.dev", "github.com", "example.com", "pkg.go.dev",
		"www.rust-lang.org", "www.python.org", "nodejs.org", "no-such-site.invalid",
	}
	start := time.Now()

	jobs := make(chan string)
	results := make(chan Result)

	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, results)
		}()
	}

	// Send the jobs, then close the channel so the workers stop.
	go func() {
		for _, site := range sites {
			jobs <- site
		}
		close(jobs)
	}()

	// Close results once every worker has finished.
	go func() {
		wg.Wait()
		close(results)
	}()

	var report []Result
	for r := range results {
		report = append(report, r)
	}
	slices.SortFunc(report, func(a, b Result) int {
		return strings.Compare(a.Site, b.Site)
	})

	for _, r := range report {
		fmt.Printf("%-21s %-7s %v\n", r.Site, r.Status, r.Took)
	}
	fmt.Println("checked", len(report), "sites in", time.Since(start).Round(time.Millisecond))
}
