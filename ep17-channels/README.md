# Episode 17: Channels

Module 4 · Concurrency

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Sending and receiving on a channel
- Unbuffered vs buffered channels
- Closing channels and ranging over them
- Directional channels (`chan<-`, `<-chan`)
- Deadlocks: what causes them and how to read the error

## Code

```bash
go run .            # all the examples from the video
go run ./solution   # the exercise solution
```

| Topic | Example |
| --- | --- |
| Create a channel | `ch := make(chan string)` |
| Send | `ch <- "hello"` |
| Receive | `msg := <-ch` |
| Buffered channel | `make(chan int, 2)` |
| Close (sender only) | `close(ch)` |
| Receive until closed | `for n := range ch { ... }` |
| Is it closed? | `v, ok := <-ch` |
| Send-only parameter | `out chan<- int` |
| Receive-only parameter | `in <-chan int` |

## Watch out

- On an unbuffered channel, a send waits for a receive, and a receive waits for a send.
- `fatal error: all goroutines are asleep - deadlock!` means every goroutine is waiting. Look for `[chan send]` or `[chan receive]` in the output.
- `panic: send on closed channel`: only the sender should close a channel.

## Exercise

Build a pipeline: one goroutine generates numbers 1–100, a second squares them, and `main` adds up the results.

The sum of the squares of 1 to 100 is `338350`.
Solution: [`solution/main.go`](./solution/main.go).
