// A web server using only Go's standard library.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// home answers GET /.
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go!")
}

// hello answers GET /hello/{name}, with an optional ?lang=es.
func hello(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	greeting := "Hello"
	if r.URL.Query().Get("lang") == "es" {
		greeting = "Hola"
	}
	fmt.Fprintf(w, "%s, %s!\n", greeting, name)
}

// logging is middleware: it wraps a handler, and logs every request.
func logging(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	}
	return http.HandlerFunc(fn)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /hello/{name}", hello)

	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", logging(mux)))
}
