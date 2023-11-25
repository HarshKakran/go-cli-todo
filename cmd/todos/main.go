package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/HarshKakran/go-cli-todo"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.String("add", "", "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	del := flag.Int("del", 0, "delete the todo")
	delAll := flag.Bool("delAll", false, "delete all the tasks")
	list := flag.Bool("list", false, "list all toods")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintf(os.Stderr, "unable to load todos. %v", err)
		os.Exit(1)
	}

	switch {
	case *add != "":
		todos.Add(*add)
		err := todos.Save(todoFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to save todos. %v", err)
			os.Exit(1)
		}

	case *complete > 0:
		err := todos.Complete(*complete)
		handleTaskError(err)
		err = todos.Save(todoFile)
		handleTaskError(err)
	case *del > 0:
		err := todos.Delete(*del)
		handleTaskError(err)
		err = todos.Save(todoFile)
		handleTaskError(err)
	case *delAll:
		todos = &todo.Todos{}
		err := todos.Save(todoFile)
		handleTaskError(err)
	case *list:
		todos.Print()
	default:
		fmt.Fprintf(os.Stdout, "invalid command")
		os.Exit(0)
	}
}

func handleTaskError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to save todos. %v", err)
		os.Exit(1)
	}
}
