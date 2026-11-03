# Episode 8: Structs & Methods

Module 2 · Working with data

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Defining and creating structs
- Struct literals and zero-value structs
- Methods on a type
- Value receivers vs pointer receivers
- Embedding structs

## Code

```bash
go run .            # all the examples from the video
go run ./solution   # the exercise solution
```

| Topic | Example |
| --- | --- |
| Define a struct | `type Person struct { Name string; Age int }` |
| Create one | `p := Person{Name: "Alex", Age: 30}` |
| Read or change a field | `p.Name`, `p.Age = 31` |
| Print with field names | `fmt.Printf("%+v\n", p)` |
| Zero value | `var empty Person` gives `{Name: Age:0}` |
| Slice of structs | `[]Person{{Name: "Alex", Age: 30}, {Name: "Sam", Age: 25}}` |
| Method (value receiver) | `func (r Rect) Area() float64` |
| Method that changes the struct | `func (r *Rect) Scale(f float64)` |
| Embedding | `type Dog struct { Animal; Breed string }` |

## Watch out

A value receiver works on a **copy**. `ScaleBroken` in `main.go` changes the copy, so the
caller's rectangle stays `{3 4}`. If a method needs to change the struct, use a pointer
receiver (`*Rect`).

## Exercise

Create a `BankAccount` struct with `Deposit`, `Withdraw` and `Balance` methods; `Withdraw` must refuse to go below zero.

Withdraw should return an error instead of going below zero.
Solution: [`solution/main.go`](./solution/main.go).
