# Episode 7: Maps

Module 2 · Working with data

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Creating maps with literals and `make`
- Adding, updating and deleting keys
- The comma-ok idiom: `value, ok := m[key]`
- Looping over a map, and why the order is random
- Sorting map keys for stable output

## Code

```bash
go run .            # all the examples from the video
go run ./solution   # the exercise solution
```

| Topic | Example |
| --- | --- |
| Create a map | `ages := map[string]int{"Alice": 31, "Bob": 25}` |
| Read a value | `ages["Alice"]` |
| Add or update | `ages["Carol"] = 40` |
| Delete | `delete(ages, "Alice")` |
| Number of keys | `len(ages)` |
| Is the key there? | `age, ok := ages["Dave"]` |
| Create an empty map | `m := make(map[string]int)` |
| Loop (random order!) | `for k, v := range ages { ... }` |
| Loop in sorted order | `for _, k := range slices.Sorted(maps.Keys(ages)) { ... }` |
| Count things | `counts[word]++` |

## Watch out

- A missing key gives the zero value (`0`, `""`, `false`). Use `v, ok := m[key]` to tell "missing" apart from "zero".
- `var m map[string]int` is a **nil** map. Reading is fine, but writing panics: `assignment to entry in nil map`. Create maps with `make` or a literal.
- Map order is random on purpose. Sort the keys if you need a fixed order.

## Exercise

Count how many times each word appears in a sentence and print the words sorted by count.

Expected output for `"go is fast and go is fun go"`: `go 3`, `is 2`, then `and 1`, `fast 1`, `fun 1`.
Hint: `words := slices.Collect(maps.Keys(counts))`, then `slices.SortFunc` with `cmp.Compare`.
Solution: [`solution/main.go`](./solution/main.go).
