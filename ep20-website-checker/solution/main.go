// Solution to the episode 20 exercise: read the sites from a file, and choose the pool size.
// Run it with: go run ./solution -workers 5 solution/sites.txt
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// readSites returns the non-empty lines of a file.
func readSites(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sites []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line := strings.TrimSpace(scanner.Text()); line != "" {
			sites = append(sites, line)
		}
	}
	return sites, scanner.Err()
}

func check(site string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://"+site, nil)
	if err != nil {
		return "bad URL"
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "DOWN"
	}
	defer resp.Body.Close()
	return resp.Status
}

func main() {
	workers := flag.Int("workers", 3, "how many sites to check at once")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("usage: checker -workers 5 sites.txt")
		os.Exit(1)
	}
	sites, err := readSites(flag.Arg(0))
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	start := time.Now()
	jobs := make(chan string)
	var wg sync.WaitGroup
	for range *workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for site := range jobs {
				fmt.Printf("%-21s %s\n", site, check(site))
			}
		}()
	}
	for _, site := range sites {
		jobs <- site
	}
	close(jobs)
	wg.Wait()
	fmt.Println("workers:", *workers, "time:", time.Since(start).Round(time.Millisecond))
}
