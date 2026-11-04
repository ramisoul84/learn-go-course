# Episode 18: select & Timeouts

Module 4 · Concurrency

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- The `select` statement
- Timeouts with `time.After`
- Non-blocking operations with `default`
- `context.Context`: cancellation and deadlines
- Passing context through your functions

## Code

```bash
go run .            # all the examples from the video
go run ./solution   # the exercise solution: uses a slow local test server
```

| Topic | Example |
| --- | --- |
| Wait for the first ready channel | `select { case msg := <-a: ... case msg := <-b: ... }` |
| Timeout | `case <-time.After(500 * time.Millisecond):` |
| Don't wait at all | `default:` |
| Deadline | `ctx, cancel := context.WithTimeout(context.Background(), d)` |
| Cancel by hand | `ctx, cancel := context.WithCancel(context.Background())` |
| Notice the end | `case <-ctx.Done():` |
| Why it ended | `ctx.Err()` |
| Check the reason | `errors.Is(err, context.DeadlineExceeded)` |
| HTTP with a context | `http.NewRequestWithContext(ctx, http.MethodGet, url, nil)` |

## Context rules

1. `ctx` is the first parameter of a function, and it's called `ctx`.
2. `defer cancel()` right after creating a context.
3. Pass it down to every slow call. Don't store it in a struct.

## Exercise

Make an HTTP request that gives up after 2 seconds using `context.WithTimeout`.

Solution: [`solution/main.go`](./solution/main.go).
