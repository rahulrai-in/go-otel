package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func executeChoice(choice int) {
	switch choice {
	case 1:
		addTask()
	case 2:
		listTasks()
	case 3:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
	}
}

func addTask() {
	fmt.Print("Enter task title: ")
	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	title = strings.TrimSuffix(title, "\n")
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	idCounter++
	task := Task{
		ID:    idCounter,
		Title: title,
	}
	tasks = append(tasks, task)
	fmt.Println("Task added successfully!")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("%d. %s\n", task.ID, task.Title)
	}
}
