package main

import (
	"fmt"
	"time"
)

type Todo struct {
	ID        int64
	Title     string
	Done      bool
	Priority  int
	CreatedAt int64
}

const (
	PriorityLow = iota
	PriorityMedium
	PriorityHigh
)

func AddTodo(todos []Todo, title string, priority int) ([]Todo, Todo) {
	t := Todo{
		ID:        time.Now().UnixNano(),
		Title:     title,
		Done:      false,
		Priority:  priority,
		CreatedAt: time.Now().Unix(),
	}
	todos = append(todos, t)
	return todos, t
}

func FindTodoByID(todos []Todo, id int64) (Todo, int, bool) {
	for i, t := range todos {
		if t.ID == id {
			return t, i, true
		}
	}
	return Todo{}, -1, false
}

func main() {
	var todos []Todo // nil slice to start

	todos, t1 := AddTodo(todos, "Buy milk", PriorityMedium)
	todos, t2 := AddTodo(todos, "Walk dog", PriorityLow)

	fmt.Println("All todos:")
	for _, t := range todos {
		fmt.Printf("ID:%d Title:%q Done:%t Priority:%d\n", t.ID, t.Title, t.Done, t.Priority)
	}

	// Find a todo
	found, idx, ok := FindTodoByID(todos, t1.ID)
	if ok {
		fmt.Println("Found:", found.Title, "at index", idx)
	} else {
		fmt.Println("Not found")
	}

	// Use t2 to avoid compilation error
	fmt.Println("Second todo:", t2.Title)
}
