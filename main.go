package main

import (
	"flag"
	"fmt"
	"os"
)

const todoFile = ".saved_todos.json"

func errorHandling(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func main() {

	add := flag.String("add", "", "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	delete := flag.Int("delete", 0, "delete a todo from the table")

	flag.Parse()

	todos := &Todos{}

	err := todos.LoadFromFile(todoFile)
	errorHandling(err)

	switch {

	case *add != "":
		todos.Add(*add)
		err := todos.SaveToFile(todoFile)
		errorHandling(err)

	case *complete > 0:
		err = todos.Complete(*complete)
		errorHandling(err)
		err = todos.SaveToFile(todoFile)
		errorHandling(err)

	case *delete > 0:
		err = todos.Delete(*delete)
		errorHandling(err)
		err = todos.SaveToFile(todoFile)
		errorHandling(err)

	default:
		fmt.Fprintln(os.Stderr, "invalid command")
		os.Exit(0)
	}
}
