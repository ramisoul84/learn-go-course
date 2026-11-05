# Episode 22: JSON REST API

Module 5 · Final project: a REST API

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Designing REST endpoints
- Decoding JSON requests and encoding JSON responses
- CRUD handlers for to-do items
- Returning the right HTTP status codes
- Testing the API with curl

## Run it

```bash
go run .

# in another terminal:
API=localhost:8080/todos
curl $API -d '{"title":"Milk"}'
curl $API -d '{"title":"Eggs"}'
curl $API
curl $API/1
curl -X PUT $API/1 -d '{"title":"Milk","done":true}'
curl -i -X DELETE $API/2
```

| Method and path | What it does | Success |
| --- | --- | --- |
| `GET /todos` | List every to-do | `200 OK` |
| `POST /todos` | Create a to-do | `201 Created` |
| `GET /todos/{id}` | Read one to-do | `200 OK` |
| `PUT /todos/{id}` | Replace a to-do | `200 OK` |
| `DELETE /todos/{id}` | Delete a to-do | `204 No Content` |

Bad JSON returns `400 Bad Request`, and an unknown id returns `404 Not Found`. Errors are JSON too: `{"error":"todo not found"}`.

The to-dos live in memory, so they disappear when the server stops. Episode 23 fixes that with a database.

## Exercise

Add input validation so a to-do with an empty title returns `400 Bad Request` with a JSON error message.

Solution: [`solution/main.go`](./solution/main.go).
