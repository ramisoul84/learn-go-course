# Episode 12: Error Handling

Module 3 · Writing idiomatic Go

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Why Go returns errors instead of throwing exceptions
- Creating errors with `errors.New` and `fmt.Errorf`
- Wrapping errors with `%w`
- Checking errors with `errors.Is` and `errors.As`
- Custom error types, and when to use `panic`

## Code

```bash
go run .            # all the examples from the video
go run ./solution   # the exercise solution
```

| Topic | Example |
| --- | --- |
| `error` is an interface | `type error interface { Error() string }` |
| A fixed message | `errors.New("disk full")` |
| A message with values | `fmt.Errorf("cannot open %s", name)` |
| Add context, keep the cause | `fmt.Errorf("read config: %w", err)` |
| Sentinel error | `var ErrNotFound = errors.New("not found")` |
| Is this error anywhere in the chain? | `errors.Is(err, ErrNotFound)` (not `==`) |
| Custom error type | `func (e *ValidationError) Error() string` |
| Get a custom error back out | `var target *ValidationError; errors.As(err, &target)` |

Wrapping turns errors into a trail of breadcrumbs:

```text
start: read config: open config.json: no such file or directory
```

## Rules

1. **Return errors** for things that can go wrong in normal use.
2. **Add context** with `%w` when you pass an error up. Never silently ignore one.
3. **`panic` only for bugs**: situations that should be impossible.

Error messages start with a lowercase letter and have no full stop, because they get joined together.

## Exercise

Write a `ParseAge(s string) (int, error)` function that returns a different, checkable error for "not a number" and "out of range".

Use two sentinel errors, `ErrNotNumber` and `ErrOutOfRange`, wrapped with `%w`, and check them with `errors.Is`.
Solution: [`solution/main.go`](./solution/main.go).
