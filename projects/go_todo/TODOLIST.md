# CLI To-Do List

Goal: Build a command-line to-do list manager.

Features:
Add, remove, and update tasks.
Mark tasks as complete or incomplete.
Persist data using a JSON file.

Learning: File I/O, struct management, basic command-line arguments.

Create a new directory for your project:
```bash
mkdir go_todo
cd go_todo
```

Initialize the Go module
```bash
go mod init go-todo
```

To add a task:
```bash
go run main.go add "Learn Go Lang Basics"
```

To remove a task:
```bash
go run main.go remove 1
```

To mark a task as complete:
```bash
go run main.go complete 1
```

To list all tasks:
```bash
go run main.go list
```