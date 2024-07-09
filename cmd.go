package main

import (
	"fmt"
)

type App struct {
	tasks []Task
}

func NewApp() *App {
	return &App{
		tasks: []Task{},
	}
}

func (app *App) Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: %s [command]", args[0])
	}
	command := args[1]

	switch command {
	case "add":
		return app.handleAdd(args[2:])
	case "list":
		return app.handleList(args[2:])
	case "delete":
		return app.handleDelete(args[2:])
	default:
		return fmt.Errorf("unknown command: %s", command)
	}
}

func (app *App) handleAdd(args []string) error {
	fmt.Println("Add", args)
	return nil
}

func (app *App) handleList(args []string) error {
	fmt.Println("List", args)
	return nil
}

func (app *App) handleDelete(args []string) error {
	fmt.Println("Delete", args)
	return nil
}
