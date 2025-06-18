package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// TaskCLI is a simple command-line interface for managing tasks.

type Task struct {
	Title string
	Done  bool
}

var tasks []Task

// addTask adds a new task to the list of tasks.
// It takes the title of the task as an argument and appends it to the tasks slice

func addTask(title string) {

	newTask := Task{
		Title: title,
		Done:  false,
	}
	tasks = append(tasks, newTask)
	fmt.Printf("Task added: %s\n", title)

}

func saveTasks() {

	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	// Save the tasks to a file named tasks.json
	os.WriteFile("tasks.json", jsonData, 0644)
	fmt.Println("Tasks saved to tasks.json")

}

func loadTasks() {
	jsonData, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("No existing tasks found, starting fresh.")
		return
	}

	err = json.Unmarshal(jsonData, &tasks)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	fmt.Println("Tasks loaded successfully.")
}

func main() {

	fmt.Println("Welcome to TaskCLI!")

	// Load existing tasks from a file at the start of the program
	loadTasks()

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command. Available commands: add, list, remove")

		return
	}

	// Check the first argument to determine the command

	switch os.Args[1] {

	case "add":

		if len(os.Args) < 3 {
			fmt.Println("Please provide a task to add.")

			return
		}

		// Join the remaining arguments to form the task title
		// This allows for task titles with spaces

		taskTitle := strings.Join(os.Args[2:], " ")
		fmt.Printf("Adding task: %s\n", taskTitle)
		addTask(taskTitle)
		saveTasks()

	case "list":
		for _, task := range tasks {
			status := " "
			if task.Done {
				status = "x"
			}
			fmt.Printf("[%s] %s\n", status, task.Title)
		}

		fmt.Println("Listing all tasks...")

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}

}
