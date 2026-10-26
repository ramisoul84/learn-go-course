# Episode 2: Variables, Types & Constants

Module 1 · Basics

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Declaring variables with `var` and the short form `:=`
- Go's basic types: int, float64, string, bool
- Zero values: why Go variables are never "undefined"
- Constants and `iota`
- Type conversion, and `fmt.Println` vs `fmt.Printf`

## Code

```bash
go run .            # all the examples from the video
go run ./solution   # the exercise solution
```

| Topic | Example |
| --- | --- |
| Long form | `var age int`, `var city string = "Berlin"`, `var score = 9.5` |
| Short form (inside functions only) | `name := "Alex"` |
| Zero values | `0`, `0`, `""`, `false` |
| Constants | `const Pi = 3.14159`, `iota` |
| Conversion | `float64(age)`, `int(height)` (cuts off decimals) |
| Printf verbs | `%d` int, `%.2f` float, `%s` string, `%v` any, `%T` type |

## Exercise

Write a program that converts a temperature from Celsius to Fahrenheit and prints both with 1 decimal place.

Formula: `f := c*9/5 + 32`. Make `c` a float64 so the division keeps the decimals.
Solution: [`solution/main.go`](./solution/main.go).
