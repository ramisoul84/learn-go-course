# Episode 13: Packages & Modules

Module 3 · Writing idiomatic Go

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Creating your own packages
- Exported (capital letter) vs unexported names
- A simple project layout, and the `internal` folder
- Adding third-party dependencies with `go get`
- What `go.mod` and `go.sum` are for

## Code

```bash
go mod tidy         # downloads github.com/google/uuid
go run .            # Hello, Alex! and a unique ID
go build -o todo ./solution && ./todo add Learn Go   # the exercise solution
```

```text
ep13-packages-modules/
├── go.mod                 module example.com/greet
├── go.sum                 checksums of downloaded code
├── main.go                package main
├── greeting/
│   └── greeting.go        package greeting
└── solution/
    ├── main.go            the to-do app, now short
    └── todo/todo.go       package todo
```

| Topic | Example |
| --- | --- |
| Name a package | `package greeting` (same as the folder) |
| Exported (usable outside) | `func Hello(...)` (capital letter) |
| Unexported (private) | `func capitalize(...)` (small letter) |
| Import your own package | `import "example.com/greet/greeting"` |
| Use it | `greeting.Hello("alex")` |
| Add missing dependencies | `go mod tidy` |
| Add a specific version | `go get github.com/google/uuid@v1.6.0` |

- A folder named `internal/` can only be imported by code inside your own module.
- Commit both `go.mod` and `go.sum`.

## Exercise

Move the to-do logic from episode 10 into its own `todo` package and keep `main` short.

Solution: [`solution/todo/todo.go`](./solution/todo/todo.go) and [`solution/main.go`](./solution/main.go).
