# Episode 23: Adding a Database (SQLite)

Module 5 · Final project: a REST API

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- How `database/sql` and drivers fit together
- Creating tables and running queries
- Prepared statements and avoiding SQL injection
- Mapping rows to structs
- Keeping the database code behind an interface

## Run it

```bash
go run .

# in another terminal:
curl localhost:8080/todos -d '{"title":"Milk"}'
curl localhost:8080/todos
```

Stop the server, start it again, and your to-dos are still there. They live in `todos.db`, next to the code.

`store.go` holds all the database code, behind the `Store` interface. `main.go` holds the handlers, and they only know the interface.

| Topic | Example |
| --- | --- |
| Add the driver | `go get modernc.org/sqlite` |
| Register the driver | `_ "modernc.org/sqlite"` |
| Open the database | `sql.Open("sqlite", "todos.db")` |
| Run a statement | `db.Exec("DELETE FROM todos WHERE id = ?", id)` |
| Read many rows | `db.Query(...)`, `rows.Next()`, `rows.Scan(...)` |
| Read one row | `db.QueryRow(...).Scan(...)` |
| No row found | `errors.Is(err, sql.ErrNoRows)` |

Always pass user input through `?` placeholders. Never join it into the SQL text.

`evil.json` is the attack from the video. Send it to your own server with `curl localhost:8080/todos -d @evil.json`, and it's saved as a harmless title. Only try this on apps you own.

`modernc.org/sqlite` is written in pure Go, so you don't need a C compiler.

## Exercise

Add a `done` filter: `GET /todos?done=true` returns only completed items.

Solution: [`solution/`](./solution/). It adds `Filter` to the store, and uses it in the `list` handler.
