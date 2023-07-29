package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	connectDb()
	defer Db.Close()

	if len(os.Args) < 2 {
		fmt.Println("Usage: task-tracker-app <command>")
		fmt.Println("Available commands: add, view, complete, delete")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-tracker-app add <description>")
			return
		}
		addTask()

	case "view":
		viewTasks()

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-tracker-app complete <task_id>")
			return
		}
		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid taskID")
			return
		}
		completeTask(taskID)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-tracker-app update <task_id> <new_description>")
			return
		}
		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		newDescription := os.Args[3]
		updateTask(taskID, newDescription)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-tracker-app delete <task_id>")
			return
		}
		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid taskID")
			return
		}
		deleteTask(taskID)

	default:
		fmt.Println("Command does not exist")
		fmt.Println("Availabe commands: add, view, complete, update, delete ")
	}
}
