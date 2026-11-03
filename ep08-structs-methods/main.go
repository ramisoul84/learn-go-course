package main

import "fmt"

// Person groups a name and an age into one type.
type Person struct {
	Name string
	Age  int
}

// Rect is a rectangle. Width and Height share a type, so they share a line.
type Rect struct {
	Width, Height float64
}

// Area has a value receiver: it gets a copy of the Rect, and only reads it.
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

// Perimeter only reads too.
func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ScaleBroken changes a copy, so the caller's Rect never changes.
func (r Rect) ScaleBroken(f float64) {
	r.Width *= f
	r.Height *= f
}

// Scale has a pointer receiver: it changes the caller's Rect.
func (r *Rect) Scale(f float64) {
	r.Width *= f
	r.Height *= f
}

// Animal and Dog show embedding: a Dog has everything an Animal has.
type Animal struct{ Name string }

func (a Animal) Describe() string {
	return "I am " + a.Name
}

type Dog struct {
	Animal
	Breed string
}

func main() {
	// The program from the start of the episode.
	r := Rect{Width: 3, Height: 4}
	fmt.Println("Rectangle", r.Width, "x", r.Height)
	fmt.Println("Area:", r.Area())
	fmt.Println("Perimeter:", r.Perimeter())
	r.Scale(2)
	fmt.Println("After Scale(2):", r.Width, "x", r.Height)

	// Creating structs.
	p := Person{Name: "Alex", Age: 30}
	fmt.Println(p.Name, p.Age)
	p.Age = 31
	fmt.Println(p)
	fmt.Printf("%+v\n", p)

	var empty Person
	fmt.Printf("%+v\n", empty)

	people := []Person{
		{Name: "Alex", Age: 30},
		{Name: "Sam", Age: 25},
	}
	for _, person := range people {
		fmt.Println(person.Name, "is", person.Age)
	}

	// Value receiver vs pointer receiver.
	box := Rect{Width: 3, Height: 4}
	box.ScaleBroken(2)
	fmt.Println(box)
	box.Scale(2)
	fmt.Println(box)

	// Embedding.
	d := Dog{
		Animal: Animal{Name: "Rex"},
		Breed:  "Beagle",
	}
	fmt.Println(d.Name, d.Breed, d.Describe())
}
