// A command-line to-do list that saves its tasks to tasks.json.
//
//	go build -o todo
//	./todo add Buy milk
//	./todo list
//	./todo done 1
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

// Task is one item on the list. The tags set the names used in the JSON file.
type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const file = "tasks.json"

// load reads the tasks from the file. No file yet means no tasks.
func load() ([]Task, error) {
	data, err := os.ReadFile(file)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

// save writes the tasks to the file as indented JSON.
func save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: todo add | list | done")
		return
	}
	tasks, err := load()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	cmd, rest := os.Args[1], os.Args[2:]
	switch cmd {
	case "add":
		title := strings.Join(rest, " ")
		tasks = append(tasks, Task{Title: title})
		fmt.Println("Added:", title)
	case "list":
		for i, t := range tasks {
			mark := " "
			if t.Done {
				mark = "x"
			}
			fmt.Printf("%d. [%s] %s\n", i+1, mark, t.Title)
		}
		return
	case "done":
		if len(rest) != 1 {
			fmt.Println("usage: todo done <n>")
			return
		}
		n, err := strconv.Atoi(rest[0])
		if err != nil || n < 1 || n > len(tasks) {
			fmt.Println("No such task.")
			return
		}
		tasks[n-1].Done = true
		fmt.Println("Done:", tasks[n-1].Title)
	default:
		fmt.Println("Unknown command:", cmd)
		return
	}

	if err := save(tasks); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
