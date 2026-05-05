package cmd

import (
	"fmt"
	"os"
)

type Command func(args []string)

var commands = map[string]Command{
	"greet":   Greet,
	"calc":    Calc,
	"upper":   Upper,
	"version": func(args []string) { Version() },
}

func Execute() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Hello, World!")
		return
	}

	cmd, ok := commands[args[0]]
	if !ok {
		if args[0] == "help" {
			PrintHelp()
			return
		}
		fmt.Printf("unknown command: %s\n\n", args[0])
		PrintHelp()
		os.Exit(1)
	}

	cmd(args[1:])
}

func PrintHelp() {
	fmt.Print(helpText)
}