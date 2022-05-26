package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aerlaut/go-todo/todo"
)

// Command names
const (
	HELP   string = "help"
	LIST   string = "list"
	DELETE string = "delete"
	ADD    string = "add"
	EXIT   string = "exit"
	SAVE   string = "save"
)

var filePath string

func executeCommand(command string, argument string, todos *[]todo.Todo) {

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

		newTodoList := make([]todo.Todo, 0)

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

		*todos = append(*todos, todo.TodoFactory(argument))
		fmt.Println("Todo added")

	case SAVE:
		saveTodos(todos, filePath)
		fmt.Println("File saved!")

	default:
		fmt.Println("Invalid command")
	}

}

func saveTodos(todos *[]todo.Todo, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file")
		return
	}

	content, _ := json.Marshal(todos)
	fmt.Fprintln(file, string(content))

	defer file.Close()
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
	filePath = *flag.String("f", "", "File to use for storing data")

	flag.Parse()

	if *showHelp {
		printHelp()
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	todos := make([]todo.Todo, 0)

	fmt.Println("Enter command: <help> to show help")
	for scanner.Scan() {

		input := scanner.Text()
		command, argument := parseCommand(input)

		if command == EXIT {
			saveTodos(&todos, filePath)
			break
		}

		executeCommand(command, argument, &todos)

		fmt.Println()
		fmt.Println("Enter command: <help> to show help")
	}

	fmt.Println("Exiting todos. Thanks you for using the program")

}
