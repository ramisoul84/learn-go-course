# Episode 20: Mini-Project: Concurrent Website Checker

Module 4 · Concurrency

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- The worker pool pattern
- Sending jobs and collecting results over channels
- Limiting how many requests run at once
- Per-request timeouts with context
- Printing a tidy results report

## Run it

Both programs need an internet connection.

```bash
go run .                                       # checks 8 sites with 3 workers
go run ./solution -workers 5 solution/sites.txt
```

```text
example.com           200 OK  279ms
github.com            200 OK  289ms
no-such-site.invalid  DOWN    72ms
...
checked 8 sites in 1.107s
```

## How it works

| Part | Job |
| --- | --- |
| `jobs` channel | Holds the sites waiting to be checked |
| `worker` goroutines | A fixed number, each taking one site at a time |
| `results` channel | Carries each `Result` back to `main` |
| `check` | Requests one site, with its own 3-second deadline |
| `close(jobs)` | Tells the workers there's no more work |
| `wg.Wait()`, then `close(results)` | Ends the `range` loop in `main` |

The sending and the waiting happen in their own goroutines, so `main` is free to receive. Without them, the program would deadlock.

Timings on a live network jump around, and one slow site can add seconds.

## Exercise

Read the list of websites from a file and add a `-workers` flag to choose the pool size.

Solution: [`solution/main.go`](./solution/main.go).
