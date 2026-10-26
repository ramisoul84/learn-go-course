// Solution to the episode 2 exercise. Run it with: go run ./solution
package main

import "fmt"

func main() {
	c := 21.5 // a float64, so the division below keeps the decimals
	f := c*9/5 + 32

	fmt.Printf("%.1f°C is %.1f°F\n", c, f)
}
