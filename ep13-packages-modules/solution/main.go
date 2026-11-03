// Solution to the episode 13 exercise: the to-do list from episode 10, with its logic
// moved into the todo package. Build it with: go build -o todo ./solution
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/greet/solution/todo"
)

func main() {
	tasks, err := todo.Load("tasks.json")
	if err == nil && len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			tasks = append(tasks, todo.Task{Title: strings.Join(os.Args[2:], " ")})
		case "list":
			todo.Print(tasks)
			return
		case "done":
			n, _ := strconv.Atoi(strings.Join(os.Args[2:], ""))
			err = todo.Complete(tasks, n)
		}
	}
	if err == nil {
		err = todo.Save("tasks.json", tasks)
	}
	if err != nil {
		fmt.Println("error:", err)
	}
}
