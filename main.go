package main

import (
	"ToDoList/scanner"
	"ToDoList/todo"
)

func main() {
	todoList := todo.NewList()

	scanner := scanner.NewScanner(todoList)

	scanner.Start()
}
