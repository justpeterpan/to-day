package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("to-day is a good day to finish some work!")
	app := NewApp()

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
