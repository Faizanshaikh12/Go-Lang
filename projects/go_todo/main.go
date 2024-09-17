package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define Task Struct
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type ToDoList struct {
	Tasks []Task `json:"tasks"`
}

// Implement File I/O For Persisting Data
const dataFile = "tasks.json"

// Load tasks from the json file
func loadTasks() (*ToDoList, error) {
	file, err := os.OpenFile(dataFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	toDoList := &ToDoList{}
	err = json.NewDecoder(file).Decode(toDoList)
	if err != nil {
		// Return empty list if is empty
		return &ToDoList{}, nil
	}
	return toDoList, nil
}

// Save tasks to the JSON file
func (td *ToDoList) saveTasks() error {
	file, err := os.OpenFile(dataFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(td)
}

// Add a new task
func (td *ToDoList) addTask(description string) {
	newTask := Task{
		ID:          len(td.Tasks) + 1,
		Description: description,
		Completed:   false,
	}
	td.Tasks = append(td.Tasks, newTask)
	fmt.Printf("Add task: %s\n", description)
}

// Remove a task by ID
func (td *ToDoList) removeTask(id int) {
	for i, task := range td.Tasks {
		if task.ID == id {
			td.Tasks = append(td.Tasks[:i], td.Tasks[i+1:]...)
			fmt.Printf("Removed task with ID: %d\n", id)
			return
		}
	}
}

// Mark a task as complete
func (td *ToDoList) markTaskComplete(id int) {
	for i, task := range td.Tasks {
		if task.ID == id {
			td.Tasks[i].Completed = true
			fmt.Printf("Task with ID: %d marked as complete.\n", id)
			return
		}
	}
	fmt.Printf("Task with ID: %d not found.\n", id)
}

// Mark a task as incomplete
func (td *ToDoList) markTaskIncomplete(id int) {
	for i, task := range td.Tasks {
		if task.ID == id {
			td.Tasks[i].Completed = false
			fmt.Printf("Task with ID: %d marked as incomplete.\n", id)
			return
		}
	}
	fmt.Printf("Task with ID: %d not found.\n", id)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: add, remove, complete, incomplete, list")
		return
	}

	// Load tasks from file
	todoList, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		description := os.Args[2]
		todoList.addTask(description)
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID to remove.")
			return
		}
		id := toInt(os.Args[2])
		todoList.removeTask(id)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID to mark complete.")
			return
		}
		id := toInt(os.Args[2])
		todoList.markTaskComplete(id)
	case "incomplete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID to mark incomplete.")
			return
		}
		id := toInt(os.Args[2])
		todoList.markTaskIncomplete(id)
	case "list":
		for _, task := range todoList.Tasks {
			status := "Incomplete"
			if task.Completed {
				status = "Complete"
			}
			fmt.Printf("[%d] %s - %s\n", task.ID, task.Description, status)
		}
	default:
		fmt.Println("Unknown command. Use add, remove, complete, incomplete, list.")
	}

	// Save tasks to file
	if err := todoList.saveTasks(); err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

// Helper to convert string to int
func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}
