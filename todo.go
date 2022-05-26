package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Command names
const (
	HELP   string = "help"
	LIST   string = "list"
	DELETE string = "delete"
	ADD    string = "add"
	EXIT   string = "exit"
)

func executeCommand(command string, argument string, todos *[]string) {

	switch command {
	// Print available commands
	case HELP:
		fmt.Println("Available commands:")
		fmt.Println("  help - Show this help")
		fmt.Println("  add <todo>- Add todo")
		fmt.Println("  list - List all todos")
		fmt.Println("  delete <id> - Delete todo by id")
		fmt.Println("  exit - Exit program")

	// List all todos
	case LIST:
		fmt.Println("Todos:")
		for idx, todo := range *todos {
			fmt.Printf("%d. %s\n", idx+1, todo)
		}

	// Delete todo by id
	case DELETE:
		deleteId, err := strconv.Atoi(argument)
		if err != nil {
			fmt.Println("Invalid id")
			return
		}

		if deleteId > len(*todos) {
			fmt.Println("Invalid ID")
			return
		}

		newTodoList := make([]string, 0)

		for idx, todo := range *todos {
			if idx+1 == deleteId {
				continue
			}
			newTodoList = append(newTodoList, todo)
		}

		*todos = newTodoList

	case ADD:
		if argument == "" {
			return
		}

		*todos = append(*todos, argument)
		fmt.Println("Todo added")

	default:
		fmt.Println("Invalid command")
	}

}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  go run main.go [options]")
	fmt.Println("")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func parseCommand(input string) (command string, argument string) {
	input = strings.TrimSpace(input)
	inputs := strings.SplitN(input, " ", 2)

	command = inputs[0]
	if len(inputs) > 1 {
		argument = strings.Trim(inputs[1], " ")
	}

	return command, argument
}

func main() {

	showHelp := flag.Bool("h", false, "Show help")
	// saveFile := flag.Bool("f", "todo.todo", "File to use for storing data")

	flag.Parse()

	if *showHelp {
		printHelp()
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	todos := make([]string, 0)

	fmt.Println("Enter command: <help> to show help")
	for scanner.Scan() {

		input := scanner.Text()
		command, argument := parseCommand(input)

		executeCommand(command, argument, &todos)

		fmt.Println()
		fmt.Println("Enter command: <help> to show help")
	}

	fmt.Println("Exiting todos. Thanks you for using the program")

}
