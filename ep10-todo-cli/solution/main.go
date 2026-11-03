// Solution to the episode 10 exercise: a delete command, and the date each task was created.
// Build and run it with:
//
//	go build -o todo ./solution
//	./todo add Buy milk
//	./todo delete 1
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	Title   string    `json:"title"`
	Done    bool      `json:"done"`
	Created time.Time `json:"created"`
}

const file = "tasks.json"

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

func save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}

// taskNumber turns the argument into a valid task number, from 1 to count.
func taskNumber(rest []string, count int) (int, bool) {
	if len(rest) != 1 {
		return 0, false
	}
	n, err := strconv.Atoi(rest[0])
	if err != nil || n < 1 || n > count {
		return 0, false
	}
	return n, true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: todo add <task> | list | done <n> | delete <n>")
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
		tasks = append(tasks, Task{Title: title, Created: time.Now()})
		fmt.Println("Added:", title)
	case "list":
		for i, t := range tasks {
			mark := " "
			if t.Done {
				mark = "x"
			}
			// Go formats dates with an example: Jan 2 means "month day".
			fmt.Printf("%d. [%s] %s (added %s)\n", i+1, mark, t.Title, t.Created.Format("Jan 2"))
		}
		return
	case "done":
		n, ok := taskNumber(rest, len(tasks))
		if !ok {
			fmt.Println("No such task.")
			return
		}
		tasks[n-1].Done = true
		fmt.Println("Done:", tasks[n-1].Title)
	case "delete":
		n, ok := taskNumber(rest, len(tasks))
		if !ok {
			fmt.Println("No such task.")
			return
		}
		fmt.Println("Deleted:", tasks[n-1].Title)
		tasks = slices.Delete(tasks, n-1, n)
	default:
		fmt.Println("Unknown command:", cmd)
		return
	}

	if err := save(tasks); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
