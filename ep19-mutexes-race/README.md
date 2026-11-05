# Episode 19: Mutexes & Race Conditions

Module 4 · Concurrency

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- What a race condition is, shown live
- Finding races with `go run -race`
- Protecting shared data with `sync.Mutex`
- `sync.RWMutex` for read-heavy data
- Atomic counters with `sync/atomic`

## Code

```bash
go run ./racy         # the broken counter: a different wrong answer every time
go run -race ./racy   # the race detector finds the bug
go run .              # fixed with a mutex, and with atomic
go run -race ./solution
```

| Topic | Example |
| --- | --- |
| Find races | `go run -race .`, `go test -race ./...` |
| Lock | `var mu sync.Mutex`, `mu.Lock()`, `mu.Unlock()` |
| Many readers, one writer | `sync.RWMutex`, `RLock`, `RUnlock` |
| Simple counters | `var total atomic.Int64`, `total.Add(1)`, `total.Load()` |

## Why updates get lost

`count++` is really three steps: read, add, write. If two goroutines read the same value,
they both write the same result, and one update is lost.

## Which tool?

- `atomic` for a single number or flag.
- A mutex when several things must change together.
- Channels when you're passing data between goroutines.

The race detector only finds races in code that actually runs, so use it with your tests.

## Exercise

Build a thread-safe `Counter` type, then prove it's safe with 1,000 goroutines and `-race`.

Solution: [`solution/main.go`](./solution/main.go).
