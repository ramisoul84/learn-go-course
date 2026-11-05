# Episode 21: HTTP Server with net/http

Module 5 · Final project: a REST API

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Starting a server with `http.ListenAndServe`
- Handlers and `http.HandlerFunc`
- Routing with methods and path values in `http.ServeMux`
- Reading query parameters
- Writing middleware for logging

## Run it

```bash
go run .

# in another terminal:
curl localhost:8080
curl localhost:8080/hello/Alex
curl "localhost:8080/hello/Alex?lang=es"
curl -i localhost:8080/nope
curl -i -X POST localhost:8080/hello/Alex
```

The solution adds a health check: `go run ./solution`, then `curl localhost:8080/health`.

| Topic | Example |
| --- | --- |
| Handler | `func home(w http.ResponseWriter, r *http.Request)` |
| Write a response | `fmt.Fprintln(w, "Hello from Go!")` |
| Router | `mux := http.NewServeMux()` |
| Route with a method | `mux.HandleFunc("GET /hello/{name}", hello)` |
| Exactly `/` | `"GET /{$}"` |
| Path value | `r.PathValue("name")` |
| Query parameter | `r.URL.Query().Get("lang")` |
| Start the server | `log.Fatal(http.ListenAndServe(":8080", mux))` |
| Middleware | `func logging(next http.Handler) http.Handler` |

Method and path patterns like `"GET /hello/{name}"` need Go 1.22 or newer.

## Exercise

Add a `/health` endpoint that returns `ok` and the server's uptime.

Solution: [`solution/main.go`](./solution/main.go).
