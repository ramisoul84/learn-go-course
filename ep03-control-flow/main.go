package main

import "fmt"

func main() {
	// The program from the start of the episode.
	for n := 1; n <= 5; n++ {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}

	// if, else if, else: the first match wins.
	temp := 28
	if temp > 25 {
		fmt.Println("It's hot")
	} else if temp < 10 {
		fmt.Println("It's cold")
	} else {
		fmt.Println("It's nice")
	}

	// if with a short statement: age only exists inside the if/else.
	if age := 17; age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("minor")
	}

	// for: the classic loop, the "while" loop, and the infinite loop.
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println(n) // 128

	for {
		fmt.Println("forever?")
		break
	}

	// range over a string, and over a number (Go 1.22+).
	for i, ch := range "Go!" {
		fmt.Println(i, string(ch))
	}
	for i := range 3 {
		fmt.Println("lap", i)
	}

	// continue skips to the next round, break leaves the loop.
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 7 {
			break
		}
		fmt.Println(i)
	}

	// switch: no break needed, several values per case, default.
	day := "Sat"
	switch day {
	case "Sat", "Sun":
		fmt.Println("Weekend!")
	case "Fri":
		fmt.Println("Almost there")
	default:
		fmt.Println("Workday")
	}

	// switch with no value: each case is a condition.
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
