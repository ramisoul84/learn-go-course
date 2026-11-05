// Solution to the episode 21 exercise. Run it with: go run ./solution
// Then try: curl localhost:8080/health
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	started := time.Now()

	mux := http.NewServeMux()
	health := func(w http.ResponseWriter, r *http.Request) {
		uptime := time.Since(started).Round(time.Second)
		fmt.Fprintf(w, "ok, up for %v\n", uptime)
	}
	mux.HandleFunc("GET /health", health)

	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
