# Episode 16: Goroutines

Module 4 · Concurrency

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Concurrency vs parallelism
- Starting a goroutine with the `go` keyword
- Why `main` doesn't wait, and how to fix it
- Waiting for goroutines with `sync.WaitGroup`
- How cheap goroutines really are

## Code

```bash
go run .            # one by one, at the same time, then 100,000 goroutines
go run ./solution   # the exercise solution: needs an internet connection
```

| Topic | Example |
| --- | --- |
| Start a goroutine | `go work("A")` |
| Anonymous goroutine | `go func() { ... }()` |
| Create a WaitGroup | `var wg sync.WaitGroup` |
| Count a new job | `wg.Add(1)` |
| Mark a job done | `defer wg.Done()` |
| Wait for all jobs | `wg.Wait()` |
| Shortcut (Go 1.25+) | `wg.Go(func() { ... })` |
| Measure time | `start := time.Now()`, then `time.Since(start)` |

## Watch out

- `go` doesn't wait. When `main` ends, the whole program ends, even if goroutines are still running.
- Don't use `time.Sleep` to wait for goroutines. That's a guess. Use a `WaitGroup`.
- The order goroutines finish in can change on every run.

## Exercise

Download 5 web pages at the same time with goroutines and print how long the whole thing took.

Solution: [`solution/main.go`](./solution/main.go).
