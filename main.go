package main

import (
	"fmt"
	"os"
	"strconv"
)


type Task struct {
	ID  	 	int
	Title 		string
	Done 		bool
}

func main() {
	fmt.Println("Task Manager")

	tasks := []Task{}

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No command provided.")
		return
	}

	switch args[0] {
	case "add":
		tasks = addTask(tasks, args[1])
	case "list":
		listTasks(tasks)
	case "complete":
		id, _ := strconv.Atoi(args[1])
		tasks = completeTask(tasks, id)
	case "delete":
		id, _ := strconv.Atoi(args[1])
		tasks = deleteTask(tasks, id)
	default:
		fmt.Println("Unknown command.")
	}
	
}

func addTask(tasks []Task, title string) []Task {
	tasks = append(tasks, Task{ID: len(tasks) + 1, Title: title, Done: false})
	return tasks
}

func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for i, task := range tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}

		fmt.Printf("%d. [%s] %s (ID: %d)\n",
			i+1,
			status,
			task.Title,
			task.ID,
		)
	}
}

func completeTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			fmt.Printf("Task '%s' marked as completed.\n", task.Title)
		}
	}
	return tasks
}

func deleteTask(tasks []Task, id int) [] Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Task '%s' deleted.\n", task.Title)
			break
		}
	}
	return tasks
}