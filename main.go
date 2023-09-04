package main

import "fmt"

// Task represents a single to-do item
type Task struct {
	ID    int
	Title string
}

var tasks []Task
var idCounter int

func main() {
	executeVanilla := false
	operationType := func() string {
		if executeVanilla {
			return "Vanilla"
		}
		return "Instrumented"
	}()

	for {
		fmt.Printf("\nTo-Do App (%s version)\n", operationType)
		fmt.Println("-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-")
		displayMenu()
		choice := getUserChoice()

		if executeVanilla {
			// Vanilla
			executeChoice(choice)
		} else {
			// Instrumented
			executeChoiceInstrumented(choice)
		}
	}
}

func displayMenu() {
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Exit")
	fmt.Print("Enter your choice: ")
}

func getUserChoice() int {
	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return 0
	}
	return choice
}
