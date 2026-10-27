# Episode 6: Arrays, Slices & How append Really Works

Module 2 · Working with data

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Arrays vs slices, and why you'll almost always use slices
- Length vs capacity
- How `append` grows a slice
- Slicing `s[a:b]` and the shared-memory surprise
- `make`, `copy` and the `slices` package

## Code

```bash
go run .            # all the examples from the video
go run ./solution   # the exercise solution
```

| Topic | Example |
| --- | --- |
| Array: fixed size, part of the type | `var scores [3]int` |
| Index starts at 0 | `scores[0] = 90` |
| Slice: can grow | `nums := []int{10, 20, 30}` |
| Add items (always assign back) | `nums = append(nums, 40)` |
| Part of a slice | `nums[1:3]`, `nums[:2]`, `nums[2:]` |
| Length and capacity | `len(s)`, `cap(s)` |
| Create with length and capacity | `make([]int, 3, 10)` |
| Copy items | `n := copy(dst, src)` |
| Independent copy | `c := slices.Clone(b)` |
| Helpers | `slices.Sort(nums)`, `slices.Contains(nums, 20)` |

## The shared memory surprise

```go
a := []int{1, 2, 3}
b := a[:2]   // b is a view onto the same array as a
b[0] = 99
fmt.Println(a) // [99 2 3]
```

Slicing never copies. When you need independent data, use `slices.Clone` or `copy`.

## How append grows a slice

When a slice is full (`len == cap`), `append` allocates a bigger array (roughly double),
copies the items over, and returns a slice that points to the new array. That's why you
always write `s = append(s, x)`. The exact capacities depend on the Go version.

## Exercise

Write a function that removes duplicates from a `[]int` while keeping the original order.

Example: `unique([]int{3, 1, 3, 2, 1})` returns `[3 1 2]`.
Hint: build the result with `append`, and skip numbers where `slices.Contains(result, n)` is true.
Solution: [`solution/main.go`](./solution/main.go).
