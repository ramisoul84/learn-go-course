# Episode 4: Functions & Multiple Return Values

Module 1 · Basics

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Declaring functions with parameters and return types
- Returning multiple values, and the `value, err` pattern
- Named return values
- Variadic functions (`...`)
- Functions as values and anonymous functions

## Code

```bash
go run .            # all the examples from the video
go run ./solution   # the exercise solution
```

| Topic | Example |
| --- | --- |
| Declare a function | `func area(width, height float64) float64 { return width * height }` |
| Multiple return values | `func divide(a, b float64) (float64, error)` |
| The value, err pattern | `result, err := divide(10, 4)` then `if err != nil { ... }` |
| Ignore a value | `_, err = divide(1, 0)` |
| Named results + bare return | `func split(total int) (half, rest int) { ...; return }` |
| Variadic | `func sum(nums ...int) int`, called as `sum(1, 2, 3)` or `sum()` |
| Functions as values | `double := func(x int) int { return x * 2 }` |
| Function parameters | `func apply(x int, f func(int) int) int` |

## Exercise

Write a function `minMax(nums ...int) (int, int)` that returns the smallest and largest number.

Example: `lo, hi := minMax(3, 9, 1, 7)` gives `1` and `9`.
Hint: start with `lo, hi := nums[0], nums[0]`, then loop and compare.
Solution: [`solution/main.go`](./solution/main.go).
