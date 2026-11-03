package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func main() {
	// The word counter from the start of the episode.
	text := "go is fast and go is fun go"
	counts := make(map[string]int)
	for _, word := range strings.Fields(text) {
		counts[word]++ // a missing word starts at 0
	}
	for _, w := range slices.Sorted(maps.Keys(counts)) {
		fmt.Println(w, counts[w])
	}

	// Creating a map and reading a value.
	ages := map[string]int{
		"Alice": 31,
		"Bob":   25,
	}
	fmt.Println(ages["Alice"]) // 31

	// Adding, updating and deleting.
	ages["Carol"] = 40 // add
	ages["Bob"] = 26   // update
	delete(ages, "Alice")
	fmt.Println(ages, len(ages)) // map[Bob:26 Carol:40] 2

	// The comma-ok idiom: is the key really there?
	fmt.Println(ages["Dave"]) // 0: missing keys give the zero value
	age, ok := ages["Dave"]
	fmt.Println(age, ok) // 0 false
	if age, ok := ages["Bob"]; ok {
		fmt.Println("Bob is", age)
	}

	// A nil map can be read, but writing to it panics. Create maps with make.
	var empty map[string]int
	fmt.Println(empty["x"]) // 0
	m := make(map[string]int)
	m["x"] = 1
	fmt.Println(m)

	// Map order is random. Sort the keys when you need a fixed order.
	ages["Dave"] = 19
	for name, age := range ages {
		fmt.Println(name, age)
	}
	for _, n := range slices.Sorted(maps.Keys(ages)) {
		fmt.Println(n, ages[n])
	}
}
