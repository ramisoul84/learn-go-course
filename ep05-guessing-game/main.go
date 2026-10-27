// A number guessing game: the program picks a secret number and gives hints.
package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

// scanner reads what the user types. Outside a function we need var, not :=.
var scanner = bufio.NewScanner(os.Stdin)

// readNumber waits for a line of input and turns it into an int.
func readNumber() (int, error) {
	if !scanner.Scan() {
		fmt.Println("Bye!")
		os.Exit(0)
	}
	text := strings.TrimSpace(scanner.Text())
	return strconv.Atoi(text)
}

func main() {
	secret := rand.IntN(100) + 1 // 1 to 100
	fmt.Println("Guess my number from 1 to 100.")

	tries := 0
	for {
		fmt.Print("Your guess: ")
		guess, err := readNumber()
		if err != nil {
			fmt.Println("Please type a whole number.")
			continue
		}
		tries++
		switch {
		case guess < secret:
			fmt.Println("Too low!")
		case guess > secret:
			fmt.Println("Too high!")
		default:
			fmt.Printf("Correct! %d tries.\n", tries)
			return
		}
	}
}
