package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      string //in-progress done todo
	CreatedAt   string
	UpdatedAt   string
}

type TaskList struct {
	Catalog []*Task
}

func (taskList *TaskList) Save() error {
	data, err := json.MarshalIndent(taskList.Catalog, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}
func (taskList *TaskList) Load() error {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			taskList.Catalog = []*Task{}
			return nil
		}
		return err
	}
	if len(file) == 0 {
		taskList.Catalog = []*Task{}
		return nil
	}
	return json.Unmarshal(file, &taskList.Catalog)
}

func (taskList *TaskList) Add(description string) error {
	task := &Task{
		Id:          len(taskList.Catalog) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format("02.01.2006 15:04"),
	}
	taskList.Catalog = append(taskList.Catalog, task)
	fmt.Printf("Task added successfully (ID: %v)\n", task.Id)
	return taskList.Save()
}

func (taskList *TaskList) Update(id int, update string) error {
	for _, task := range taskList.Catalog {
		if task.Id == id {
			task.Description = update
			task.UpdatedAt = time.Now().Format("02.01.2006 15:04")
		}
	}
	return taskList.Save()
}

func (taskList *TaskList) Delete(id int) error {
	for number, task := range taskList.Catalog {
		if task.Id == id {
			taskList.Catalog = append(taskList.Catalog[:number], taskList.Catalog[number+1:]...)
			for i := number; i < len(taskList.Catalog); i++ {
				taskList.Catalog[i].Id--
			}
			break
		}
	}
	return taskList.Save()
}

func (taskList *TaskList) MarkInProgress(id int) error {
	for _, task := range taskList.Catalog {
		if task.Id == id {
			task.Status = "in-progress"
			task.UpdatedAt = time.Now().Format("02.01.2006 15:04")
		}
	}
	return taskList.Save()
}

func (taskList *TaskList) MarkDone(id int) error {
	for _, task := range taskList.Catalog {
		if task.Id == id {
			task.Status = "done"
			task.UpdatedAt = time.Now().Format("02.01.2006 15:04")
		}
	}
	return taskList.Save()
}

func (taskList *TaskList) ListByStatus(status string) {
	if len(taskList.Catalog) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range taskList.Catalog {
		if status == "" || task.Status == status {
			fmt.Printf("[%d] %-20s | %-11s | Created: %s\n",
				task.Id, task.Description, task.Status, task.CreatedAt)
		}
	}
}

func (taskList *TaskList) CheckArgument(argument string) int {
	i, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Invalid ID format. Please enter a number.")
		return -1
	}
	if i < 1 || i > len(taskList.Catalog) {
		fmt.Println("Task with this ID does not exist.")
		return -1
	}
	return i
}

func CheckNumberArgument(arguments []string, requiredLen int) bool {
	if len(arguments) != requiredLen {
		fmt.Println("Invalid arguments number")
		return false
	}
	return true
}

func main() {
	var taskList TaskList

	if err := taskList.Load(); err != nil {
		fmt.Println("Load mistake: ", err)
		return
	}
	if len(os.Args) == 1 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		fmt.Println("Commands: add, update, delete, mark-in-progress, mark-done, list")
		return
	}

	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "add":
			if CheckNumberArgument(os.Args, 3) {
				taskList.Add(os.Args[2])
			}
		case "update":
			if CheckNumberArgument(os.Args, 4) {
				if id := taskList.CheckArgument(os.Args[2]); id != -1 {
					taskList.Update(id, os.Args[3])
				}
			}
		case "delete":
			if CheckNumberArgument(os.Args, 3) {
				if id := taskList.CheckArgument(os.Args[2]); id != -1 {
					taskList.Delete(id)
				}
			}
		case "mark-in-progress":
			if CheckNumberArgument(os.Args, 3) {
				if id := taskList.CheckArgument(os.Args[2]); id != -1 {
					taskList.MarkInProgress(id)
				}
			}
		case "mark-done":
			if CheckNumberArgument(os.Args, 3) {
				if id := taskList.CheckArgument(os.Args[2]); id != -1 {
					taskList.MarkDone(id)
				}
			}
		case "list":
			if len(os.Args) > 2 {
				if CheckNumberArgument(os.Args, 3) {
					switch os.Args[2] {
					case "done":
						taskList.ListByStatus("done")
					case "todo":
						taskList.ListByStatus("todo")
					case "in-progress":
						taskList.ListByStatus("in-progress")
					default:
						fmt.Println("Incorrect input")
					}
				}
			} else {
				taskList.ListByStatus("")
			}
		default:
			fmt.Println("Incorrect input")
		}
	}
}
