// Solution to the episode 5 exercise: difficulty levels and a guess limit.
// Run it with: go run ./solution
package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

// readLine waits for a line of input and returns it without extra spaces.
func readLine() string {
	if !scanner.Scan() {
		fmt.Println("Bye!")
		os.Exit(0)
	}
	return strings.TrimSpace(scanner.Text())
}

// chooseLevel asks for a difficulty and returns the highest number and the guess limit.
func chooseLevel() (int, int) {
	for {
		fmt.Print("Choose a level (easy, medium, hard): ")
		switch readLine() {
		case "easy":
			return 10, 4
		case "medium":
			return 100, 7
		case "hard":
			return 1000, 10
		default:
			fmt.Println("Please type easy, medium or hard.")
		}
	}
}

func main() {
	maxNumber, maxTries := chooseLevel()
	secret := rand.IntN(maxNumber) + 1
	fmt.Printf("I'm thinking of a number between 1 and %d. You have %d tries.\n", maxNumber, maxTries)

	for tries := 1; tries <= maxTries; {
		fmt.Printf("Guess %d of %d: ", tries, maxTries)
		guess, err := strconv.Atoi(readLine())
		if err != nil {
			fmt.Println("Please type a whole number.")
			continue // a typo doesn't use up a try
		}
		switch {
		case guess < secret:
			fmt.Println("Too low!")
		case guess > secret:
			fmt.Println("Too high!")
		default:
			fmt.Printf("Correct! %d tries.\n", tries)
			return
		}
		tries++
	}
	fmt.Println("Out of tries! The number was", secret)
}
