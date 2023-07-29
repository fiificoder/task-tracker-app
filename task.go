package main

import (
	"fmt"
	"os"
	"strings"
)

func addTask() {

	description := strings.Join(os.Args[2:], " ")

	query := `
        INSERT INTO tasks (description, completed)
        VALUES ($1, $2)
        RETURNING id
    `
	var taskID int
	err := Db.QueryRow(query, description, false).Scan(&taskID)
	if err != nil {
		fmt.Println("Error adding task:", err)
		return
	}

	newTask := Task{
		ID:          taskID,
		Description: description,
		Completed:   false,
	}

	fmt.Println("Task added successfully! Task ID:", taskID)
	fmt.Println("Task:", newTask)
}

func viewTasks() {
	query := `
        SELECT id, description, completed
        FROM tasks
    `

	rows, err := Db.Query(query)
	if err != nil {
		fmt.Println("Error retrieving tasks:", err)
		return
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Description, &task.Completed)
		if err != nil {
			fmt.Println("Error scanning task:", err)
			return
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error processing tasks:", err)
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: %s, Completed: %t\n", task.ID, task.Description, task.Completed)
	}
}

func completeTask(taskID int) {
	query := `
        UPDATE tasks
        SET completed = true
        WHERE id = $1
    `

	result, err := Db.Exec(query, taskID)
	if err != nil {
		fmt.Println("Error completing task:", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting affected rows:", err)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("No task found with ID:", taskID)
	} else {
		fmt.Println("Task with ID", taskID, "marked as completed.")
	}
}

func deleteTask(taskID int) {
	query := `
        DELETE FROM tasks
        WHERE id = $1
    `

	result, err := Db.Exec(query, taskID)
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting affected rows:", err)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("No task found with ID:", taskID)
	} else {
		fmt.Println("Task with ID", taskID, "deleted.")
	}
}

func updateTask(taskID int, newDescription string) {
	query := `
        UPDATE tasks
        SET description = $1
        WHERE id = $2
    `

	result, err := Db.Exec(query, newDescription, taskID)
	if err != nil {
		fmt.Println("Error updating task:", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting affected rows:", err)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("No task found with ID:", taskID)
	} else {
		fmt.Println("Task with ID", taskID, "updated.")
	}
}
