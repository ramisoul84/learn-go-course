# Episode 19: Mutexes & Race Conditions

Module 4 · Concurrency

▶️ Watch on YouTube: [Learn Go — full course playlist](PLAYLIST_LINK)

## What you'll learn

- What a race condition is, shown live
- Finding races with `go run -race`
- Protecting shared data with `sync.Mutex`
- `sync.RWMutex` for read-heavy data
- Atomic counters with `sync/atomic`

## Code

Code for this episode will be added when the video is released.

## Exercise

Build a thread-safe `Counter` type, then prove it's safe with 1,000 goroutines and `-race`.
