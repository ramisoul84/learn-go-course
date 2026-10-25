# Learn Go — From Zero to a REST API

Code for the **Learn Go** YouTube course: 24 short, hands-on episodes that take you from
installing Go to building and deploying a REST API with a database and Docker.

▶️ **Watch the course:** [YouTube playlist](PLAYLIST_LINK)

## How to use this repo

Each episode has its own folder and its own Go module. To run an episode's code:

```bash
git clone https://github.com/YOUR_GITHUB_USERNAME/learn-go-course.git
cd learn-go-course/ep01-hello-world
go run .
```

Exercise solutions are in each episode's `solution` folder (`go run ./solution`).
Try the exercise yourself first!

You need [Go](https://go.dev/dl/) 1.24 or newer.

## Episodes

| # | Episode | Topics |
| --- | --- | --- |
| | **Module 1 · Basics** | |
| 1 | [Install Go & Write Your First Program](./ep01-hello-world) | Install Go on Windows, Mac or Linux and check it with `go version` … |
| 2 | [Variables, Types & Constants](./ep02-variables) | Declaring variables with `var` and the short form `:=` … |
| 3 | [Control Flow: if, for & switch](./ep03-control-flow) | `if` / `else` and the `if` with a short statement … |
| 4 | [Functions & Multiple Return Values](./ep04-functions) | Declaring functions with parameters and return types … |
| 5 | [Mini-Project: Number Guessing Game](./ep05-guessing-game) | Reading user input with `bufio.Scanner` … |
| | **Module 2 · Working with data** | |
| 6 | [Arrays, Slices & How append Really Works](./ep06-slices) | Arrays vs slices, and why you'll almost always use slices … |
| 7 | [Maps](./ep07-maps) | Creating maps with literals and `make` … |
| 8 | [Structs & Methods](./ep08-structs-methods) | Defining and creating structs … |
| 9 | [Pointers Without Fear](./ep09-pointers) | What a pointer is, in one picture … |
| 10 | [Mini-Project: CLI To-Do List](./ep10-todo-cli) | Command-line arguments with `os.Args` … |
| | **Module 3 · Writing idiomatic Go** | |
| 11 | [Interfaces](./ep11-interfaces) | Defining an interface … |
| 12 | [Error Handling](./ep12-errors) | Why Go returns errors instead of throwing exceptions … |
| 13 | [Packages & Modules](./ep13-packages-modules) | Creating your own packages … |
| 14 | [Testing with go test](./ep14-testing) | Writing a test function in a `_test.go` file … |
| 15 | [Generics](./ep15-generics) | Why generics were added to Go … |
| | **Module 4 · Concurrency** | |
| 16 | [Goroutines](./ep16-goroutines) | Concurrency vs parallelism … |
| 17 | [Channels](./ep17-channels) | Sending and receiving on a channel … |
| 18 | [select & Timeouts](./ep18-select-context) | The `select` statement … |
| 19 | [Mutexes & Race Conditions](./ep19-mutexes-race) | What a race condition is, shown live … |
| 20 | [Mini-Project: Concurrent Website Checker](./ep20-website-checker) | The worker pool pattern … |
| | **Module 5 · Final project: a REST API** | |
| 21 | [HTTP Server with net/http](./ep21-http-server) | Starting a server with `http.ListenAndServe` … |
| 22 | [JSON REST API](./ep22-rest-api) | Designing REST endpoints … |
| 23 | [Adding a Database (SQLite)](./ep23-database) | How `database/sql` and drivers fit together … |
| 24 | [Testing & Docker Deployment](./ep24-docker-deploy) | Testing HTTP handlers with `httptest` … |

## Questions?

Ask in the YouTube comments or open an issue in this repo.
