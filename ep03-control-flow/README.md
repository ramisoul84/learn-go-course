# Episode 3: Control Flow: if, for & switch

Module 1 · Basics

▶️ Watch on YouTube: [Learn Go — full course playlist](PLAYLIST_LINK)

## What you'll learn

- `if` / `else` and the `if` with a short statement
- `for` as a classic loop, a while loop and an infinite loop
- Looping over strings and collections with `range`
- `switch` without `break`, and `switch` with no condition
- `break` and `continue`

## Code

```bash
go run .            # all the examples from the video
go run ./solution   # the exercise solution
```

| Topic | Example |
| --- | --- |
| if / else | `if temp > 25 { ... } else if temp < 10 { ... } else { ... }` |
| if with a short statement | `if age := 17; age >= 18 { ... }` |
| Comparison and logic | `==` `!=` `<` `>` `<=` `>=`, `&&` `\|\|` `!` |
| Classic loop | `for i := 1; i <= 3; i++ { ... }` |
| "While" loop | `for n < 100 { ... }` |
| Infinite loop | `for { ... break }` |
| range | `for i, ch := range "Go!" { ... }`, `for i := range 3 { ... }` |
| Skip / stop | `continue`, `break` |
| switch | `switch day { case "Sat", "Sun": ... default: ... }` |
| switch with no value | `switch { case hour < 12: ... }` |

## Exercise

FizzBuzz: print 1 to 100, but "Fizz" for multiples of 3, "Buzz" for multiples of 5, and "FizzBuzz" for both.

Hint: use the remainder operator `%` and check for 15 first.
Solution: [`solution/main.go`](./solution/main.go).
