package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	ID        int  `json:"id"`
	Title     string  `json:"title"`
	Completed bool `json:"completed"`
}

func addTodo(title string, todos []Todo) []Todo {
	id := len(todos) + 1
	todo := Todo{ID: id, Title: title, Completed: false}
	todos = append(todos, todo)
	saveTodos(todos)
	fmt.Println("Added:", title)
	return todos
}

func listTodos(todos []Todo) {
	for _, todo := range todos {
		status := " "
		if todo.Completed {
			status = "x"
		}
		fmt.Println(status, todo.ID, todo.Title)
	}
}

func markDone(id int, todos []Todo) []Todo {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true
			saveTodos(todos)
			fmt.Println("Marked as done:", todo.Title)
			return todos
		}
	}
	fmt.Println("Todo not Found")
	return todos
}

func deleteTodo(id int, todos []Todo) []Todo {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			saveTodos(todos)
			fmt.Println("Deleted:", todo.Title)
			return todos
		}
	}
	fmt.Println("Todo not found")
	return todos
}

const filename = "todos.json"

func saveTodos(todos []Todo) {
	file, _ := json.MarshalIndent(todos, "", "	")
	_=os.WriteFile(filename, file, 0644)
}

func loadTodos() []Todo {
	var todos []Todo
	file, err := os.ReadFile(filename)
	if err == nil {
		json.Unmarshal(file, &todos)
	}
	return todos
}

func main() {
	todos := loadTodos()
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'list', 'done', or 'delete' command.")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a title.")
			return
		}
		title := os.Args[2]
		todos = addTodo(title, todos)
	
	case "list":
		listTodos(todos)
	
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide an ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2]) 
        if err != nil {
            fmt.Println("Invalid ID. Please provide a number.")
            return
        }
		todos = markDone(id, todos)
	
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide an ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2]) 
        if err != nil {
            fmt.Println("Invalid ID. Please provide a number.")
            return
        }
		todos = deleteTodo(id, todos)

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}