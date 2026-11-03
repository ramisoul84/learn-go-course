// Solution to the episode 7 exercise. Run it with: go run ./solution
package main

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"strings"
)

func main() {
	text := "go is fast and go is fun go"
	counts := make(map[string]int)
	for _, word := range strings.Fields(text) {
		counts[word]++
	}

	words := slices.Collect(maps.Keys(counts))
	slices.SortFunc(words, func(a, b string) int {
		// Highest count first. For equal counts, sort alphabetically.
		return cmp.Or(
			cmp.Compare(counts[b], counts[a]),
			cmp.Compare(a, b),
		)
	})

	for _, w := range words {
		fmt.Println(w, counts[w])
	}
}
