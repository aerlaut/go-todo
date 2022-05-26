package main

import "testing"

func Test_Help_ShouldPrintAvailableCommand(t *testing.T) {
	todos := make([]string, 0)

	ExecuteCommand(HELP, "", &todos)

	if len(todos) != 0 {
		t.Fail()
	}

}

func Test_Add_ShouldAddTodo(t *testing.T) {
	todos := make([]string, 0)

	// Add first todo
	ExecuteCommand(ADD, "test todo 1", &todos)

	if len(todos) != 1 {
		t.Fail()
	}

	// Add second todo
	ExecuteCommand(ADD, "test todo 2", &todos)
	if len(todos) != 2 {
		t.Fail()
	}
}

func Test_List_ShouldListTodos(t *testing.T) {
	todos := []string{"test todo 1", "test todo 2"}

	// Add first todo
	ExecuteCommand(LIST, "", &todos)

	if len(todos) != 2 {
		t.Fail()
	}
}

func Test_Delete_ShouldDeleteTodo(t *testing.T) {
	todos := []string{"test todo 1", "test todo 2"}

	// Add first todo
	ExecuteCommand(DELETE, "2", &todos)

	if len(todos) != 1 {
		t.Fail()
	}
}
