package main

import (
	"fmt"

	"example.com/greet/greeting"
	"github.com/google/uuid"
)

func main() {
	fmt.Println(greeting.Hello("alex"))
	fmt.Println("Your ID:", uuid.NewString())
}
