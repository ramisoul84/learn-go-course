# Episode 1: Install Go & Write Your First Program

Module 1 · Basics

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Install Go on Windows, Mac or Linux and check it with `go version`
- Set up VS Code with the official Go extension
- Create a module with `go mod init`
- What `package main`, `import` and `func main` mean
- The difference between `go run` and `go build`

## Commands from the video

```bash
go version
mkdir hello
cd hello
go mod init hello
code .
go run .
go build
./hello        # on Windows: .\hello.exe
```

## Run this folder

```bash
go run .
```

Output:

```text
Hello, World!
Welcome to Learn Go, episode 1.
```

## Exercise

Change the program so it stores your name and the language you're learning in two
variables, then prints them. Hint: `name := "your name"` creates a variable.

Solution: [`solution/main.go`](./solution/main.go), run it with `go run ./solution`.

## Common errors

| Error message | Cause | Fix |
| --- | --- | --- |
| `go: command not found` | Terminal opened before install | Open a new terminal or restart |
| `syntax error: unexpected semicolon or newline before {` | Opening brace on its own line | Put `{` on the same line as `func main()` |
| `"os" imported and not used` | Unused import | Delete it — Go refuses unused imports |
| `go: go.mod file not found in current directory` | Ran `go run .` outside the project folder | `cd hello` first |
