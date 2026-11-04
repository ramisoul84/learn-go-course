# Episode 14: Testing with go test

Module 3 · Writing idiomatic Go

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Writing a test function in a `_test.go` file
- Running tests with `go test`
- Table-driven tests
- Subtests with `t.Run`
- Test coverage with `-cover`

## Code

```bash
go test             # run the tests in this folder
go test -v          # show every test by name
go test -cover      # show how much of the code the tests ran
go test ./...       # test every package, including the solution
```

To see the failure from the video, change `score >= 90` to `score > 90` in `grade.go` and run `go test`.

| Topic | Example |
| --- | --- |
| Test file | `grade_test.go`, in the same folder and package |
| Test function | `func TestLetter(t *testing.T)` |
| Report a failure | `t.Errorf("Letter(95) = %q, want %q", got, want)` |
| Stop this test right away | `t.Fatalf(...)` |
| Table-driven test | a `[]struct{...}` of cases and one loop |
| Named subtest | `t.Run(tt.name, func(t *testing.T) { ... })` |

Test the **edges**: `Letter(95)` passed even with the bug. The case that caught it was exactly `90`.

100% coverage doesn't prove your code has no bugs. Good cases matter more than the number.

## Exercise

Write table-driven tests for the `ParseAge` function from episode 12, covering at least 5 cases.

Test the edges (`"0"`, `"150"`, `"151"`), and check errors with `errors.Is`.
Solution: [`solution/age_test.go`](./solution/age_test.go).
