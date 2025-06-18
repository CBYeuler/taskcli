# TaskCLI

A simple, beginner-friendly command-line to-do app written in Go.  
It lets you add and list tasks — with persistence using a local `tasks.json` file.

---

## Features

-  Add tasks with a simple command
-  List all tasks with completion status
-  Automatically saves to and loads from `tasks.json`
-  Handles multi-word task titles
-  Fully written in Go with no external dependencies

---

## Requirements

- Go 1.18 or higher

---

## Installation


git clone https://github.com/yourusername/taskcli.git
cd taskcli
go run main.go



## Adding a Task

go run main.go "Task"

## Listing all tasks

go run main.go list


## Project Structure

taskcli/
├── cmd/
│   └── main.go         # Entry point of the CLI app
├── task/               # (Optional) Directory for future logic and refactoring
├── tasks.json          # Auto-created JSON file to store tasks
├── go.mod              # Go module definition
├── .gitignore
└── README.md

Make sure to cd into cmd folder to run this app

This project was built to learn:
-Go fundamentals (syntax, functions, structs, slices)
-Command-line argument parsing
-File I/O and JSON persistence
-Real-World project structure

