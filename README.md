# Task Tracker CLI

A simple command-line interface (CLI) application to manage and track your tasks.  
This tool allows you to add, update, delete, and track tasks as **todo**, **in-progress**, or **done**. All tasks are stored in a JSON file in the current directory, allowing your data to persist between sessions.

---

## Assignment

This project is based on the following task: [Task Tracker Assignment](https://roadmap.sh/projects/task-tracker)  
*The assignment requires building a CLI application to track tasks, storing them in a JSON file, and handling user input via positional arguments.*

---

## Features

- Add a new task
- Update an existing task
- Delete a task
- Mark a task as **in-progress** or **done**
- List all tasks
- List tasks by status: **todo**, **in-progress**, or **done**

Each task has the following properties:

- `id`: Unique identifier
- `description`: Short description of the task
- `status`: Task status (`todo`, `in-progress`, `done`)
- `createdAt`: Date and time the task was created
- `updatedAt`: Date and time the task was last updated

---

## Installation

1. Clone this repository:
```bash
git clone https://github.com/mariapylkova/task-tracker-go.git
cd task-tracker-go
```
2. Build the application:
```bash
go build task-cli.go
```
3. Run the CLI:
```bash
./task-cli <command> [arguments]
```

## Usage

Below are example commands showing how the CLI can be used:

```bash
# Add a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Update a task
task-cli update 1 "Buy groceries and cook dinner"

# Delete a task
task-cli delete 1
# IDs of following tasks will shift (1,2,3 → delete ID=2 → becomes 1,2)

# Mark a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# List all tasks
task-cli list

# List tasks by status
task-cli list todo
task-cli list in-progress
task-cli list done
```
## Error Handling

The program gracefully handles edge cases such as:
- Invalid command usage
- Non-existent task ID
- Attempt to mark or update a deleted task

Clear error messages are shown when incorrect input is provided.

## License
This project is for educational use and may be freely improved, extended, or adapted.
