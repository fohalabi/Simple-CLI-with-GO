package main

import "fmt"


type Task struct {
	ID  	 	int
	Title 		string
	Done 		bool
}

func main() {
	fmt.Println("Task Manager")

	tasks := []Task{}
	tasks = addTask(tasks, "Buy groceries")
	listTasks(tasks)
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