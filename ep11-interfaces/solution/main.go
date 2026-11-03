// Solution to the episode 11 exercise. Run it with: go run ./solution
package main

import (
	"fmt"
	"math"
)

// Shape is anything with an area and a perimeter.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct{ Width, Height float64 }

func (r Rect) Area() float64      { return r.Width * r.Height }
func (r Rect) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Circle struct{ Radius float64 }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

func main() {
	shapes := []Shape{
		Rect{Width: 3, Height: 4},
		Circle{Radius: 1},
	}
	for _, s := range shapes {
		fmt.Printf("%T: area %.2f, perimeter %.2f\n", s, s.Area(), s.Perimeter())
	}
}
