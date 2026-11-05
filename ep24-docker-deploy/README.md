# Episode 24: Testing & Docker Deployment

Module 5 · Final project: a REST API

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Testing HTTP handlers with `httptest`
- Configuration with environment variables
- Graceful shutdown
- A multi-stage Dockerfile for a small image
- Running the API with Docker

## Run it

```bash
go test ./...               # run the tests
PORT=9000 go run .          # run it on another port

docker build -t todoapi .
docker run -d --name todos -p 8080:8080 -v todo-data:/data todoapi
curl localhost:8080/todos -d '{"title":"Milk"}'
docker stop todos
```

The `todo-data` volume keeps the database when the container is removed.

| Setting | Default | What it does |
| --- | --- | --- |
| `PORT` | `8080` | The port to listen on |
| `DB_PATH` | `todos.db` | Where the database file lives (`/data/todos.db` in Docker) |

| Topic | Example |
| --- | --- |
| Fake a request | `httptest.NewRequest(method, path, body)` |
| Record a response | `httptest.NewRecorder()` |
| Database for tests | `OpenStore(":memory:")` |
| Read a setting | `os.Getenv("PORT")` |
| Catch stop signals | `signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)` |
| Stop gently | `server.Shutdown(ctx)` |
| Small image | Build in one stage, `COPY --from=build` into the next |

## Exercise

Deploy the container to a free cloud host and share your URL in the comments!

Hints:

- Many hosts can build straight from your `Dockerfile`.
- Many hosts tell your app which port to use through `PORT`, and this app already reads it.
- Check whether your host keeps files between restarts. On many free plans, the database is wiped unless you attach a disk or volume.

There's no solution folder for this one: your live URL is the solution. Congratulations on finishing the course!
