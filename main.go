package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	showHelp := flag.Bool("h", false, "Show help")
	// saveFile := flag.Bool("f", "todo.todo", "File to use for storing data")

	flag.Parse()

	if *showHelp {
		fmt.Println("Usage:")
		fmt.Println("  go run main.go [options]")
		fmt.Println("")
		fmt.Println("Options:")
		flag.PrintDefaults()
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	todos := make([]string, 0)

	fmt.Println("Enter command: <help> to show help")
	for scanner.Scan() {

		var command, argument string

		command = scanner.Text()
		inputs := strings.SplitN(command, " ", 2)

		command = inputs[0]

		if len(inputs) > 1 {
			argument = strings.Trim(inputs[1], " ")
		}

		if command == "exit" {
			break
		}

		switch command {
		// Print available commands
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  help - Show this help")
			fmt.Println("  add <todo>- Add todo")
			fmt.Println("  list - List all todos")
			fmt.Println("  delete <id> - Delete todo by id")
			fmt.Println("  exit - Exit program")

		// List all todos
		case "list":
			fmt.Println("Todos:")
			for idx, todo := range todos {
				fmt.Printf("%d. %s\n", idx+1, todo)
			}

		// Delete todo by id
		case "delete":
			deleteId, err := strconv.Atoi(argument)
			if err != nil {
				fmt.Println("Invalid id")
				continue
			}

			if deleteId > len(todos) {
				fmt.Println("Invalid ID")
				continue
			}

			todos = append(todos[:deleteId-1], todos[deleteId:]...)

		case "add":
			if argument == "" {
				continue
			}

			todos = append(todos, argument)
			fmt.Println("Todo added")

		default:
			fmt.Println("Invalid command")
		}

		fmt.Println()
		fmt.Println("Enter command: <help> to show help")
	}

	fmt.Println("Exiting todos. Thanks you for using the program")

}
