# Episode 9: Pointers Without Fear

Module 2 · Working with data

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- What a pointer is, in one picture
- `&` to get an address, `*` to follow it
- Why Go passes everything by value
- When to use pointers (and when not to)
- `nil` pointers and how to avoid panics

## Code

```bash
go run .            # all the examples from the video
go run ./solution   # the exercise solution
```

| Symbol | Where | Read it as |
| --- | --- | --- |
| `*int` | in a type | "a pointer to an int" |
| `&x` | in front of a variable | "the address of x" |
| `*p` | in front of a pointer | "the value p points to" |
| `nil` | | "points to nothing" (the zero value of a pointer) |

## When to use a pointer

1. A function or method must **change** the caller's data (`double(&x)`, `func (a *Account) Deposit`).
2. A struct is **big**, and copying it on every call would be wasteful.
3. A value might be **missing**: a pointer can be `nil`.

Small values you only read can be passed normally. Slices and maps already share their data.

## Watch out

Following a nil pointer crashes the program:
`panic: runtime error: invalid memory address or nil pointer dereference`.
Check `if p != nil` before using a pointer that might be empty.

## Exercise

Write a `swap(a, b *int)` function, then explain in a comment why `swap(a, b int)` can't work.

Example: with `x, y := 1, 2`, after `swap(&x, &y)` x is 2 and y is 1.
Solution: [`solution/main.go`](./solution/main.go), including why `swap(a, b int)` can't work.
