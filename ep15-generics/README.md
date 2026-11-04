# Episode 15: Generics

Module 3 · Writing idiomatic Go

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Why generics were added to Go
- Type parameters on functions
- Constraints, `any`, `comparable` and `cmp.Ordered`
- Generic types, such as a `Stack[T]`
- When not to use generics

## Code

```bash
go run .            # all the examples from the video
go run ./solution   # the exercise solution
```

| Topic | Example |
| --- | --- |
| Type parameter | `func Largest[T cmp.Ordered](a, b T) T` |
| Type inference | `Largest(3, 9)` (Go works out `T`) |
| Naming the type | `Largest[float64](1, 2.5)` |
| Any type | `[T any]` |
| Works with `==` | `[T comparable]` |
| Works with `<` and `>` | `[T cmp.Ordered]` |
| Generic type | `type Stack[T any] struct { items []T }` |
| Using it | `var names Stack[string]` |

## When to use generics

- **Use them** when the logic is the same for every type, like containers and helpers.
- **Use an interface** when different types need different behaviour.
- **Start with plain code**, and reach for generics when you catch yourself copying a function.

Go has a built-in `max` (since Go 1.21), so you don't need your own `Largest`. It's here to show how generics work.

## Exercise

Write a generic `Filter[T any](items []T, keep func(T) bool) []T` and use it with ints and strings.

Solution: [`solution/main.go`](./solution/main.go).
