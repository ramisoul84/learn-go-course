# Episode 11: Interfaces

Module 3 · Writing idiomatic Go

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Defining an interface
- Implicit implementation: no `implements` keyword
- Small interfaces from the standard library: `fmt.Stringer`, `io.Reader`
- The empty interface `any` and type switches
- "Accept interfaces, return structs"

## Code

```bash
go run .            # all the examples from the video
go run ./solution   # the exercise solution
```

| Topic | Example |
| --- | --- |
| Define an interface | `type Notifier interface { Notify(msg string) string }` |
| Implement it (no `implements`!) | `func (e Email) Notify(msg string) string { ... }` |
| Use it | `func sendAll(ns []Notifier, msg string)` |
| Mix types in one slice | `[]Notifier{Email{...}, SMS{...}, Slack{...}}` |
| `fmt.Stringer` | `func (t Temp) String() string` |
| Any value | `func describe(v any) string` |
| Type switch | `switch x := v.(type) { case int: ... }` |
| Type assertion | `s, ok := v.(string)` |

## Tips

- **Accept interfaces, return structs.**
- Keep interfaces small: one or two methods.
- Define an interface where it's *used*, not next to the types that implement it.

If a type is missing a method, the compiler tells you:
`Phone does not implement Notifier (missing method Notify)`.

## Exercise

Create a `Shape` interface with `Area()` and `Perimeter()`, implement it for a circle and a rectangle, and print a list of shapes.

For the circle, use `math.Pi`. Put both shapes in a `[]Shape` and loop over it.
Solution: [`solution/main.go`](./solution/main.go).
