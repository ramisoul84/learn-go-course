# Episode 10: Mini-Project: CLI To-Do List

Module 2 · Working with data

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Command-line arguments with `os.Args`
- Modelling tasks with structs and slices
- Saving and loading JSON with `encoding/json`
- Reading and writing files with `os`
- Adding, listing and completing tasks

## Build and use it

```bash
go build -o todo
./todo add Buy milk
./todo add Learn Go
./todo list
./todo done 2
cat tasks.json
```

```text
1. [ ] Buy milk
2. [x] Learn Go
```

The exercise solution adds `delete` and a creation date:

```bash
go build -o todo ./solution
./todo delete 1
```

## How it works

| Piece | Code |
| --- | --- |
| Command-line arguments | `os.Args` (`[0]` is the program, `[1]` the command) |
| Join words into one title | `strings.Join(rest, " ")` |
| JSON field names | struct tags: `` Title string `json:"title"` `` |
| Go to JSON | `json.MarshalIndent(tasks, "", "  ")` |
| JSON to Go | `json.Unmarshal(data, &tasks)` (needs a pointer) |
| Write / read a file | `os.WriteFile(file, data, 0644)`, `os.ReadFile(file)` |
| "No file yet" is not an error | `errors.Is(err, fs.ErrNotExist)` |

JSON only sees **exported** fields (capital letters). A lowercase field is silently left out.

## Exercise

Add a `delete` command and show the date each task was created.

Hints: `slices.Delete(tasks, n-1, n)` removes a task, and a `Created time.Time` field set with `time.Now()` stores the date.
Solution: [`solution/main.go`](./solution/main.go).
