package cmd

import (
	"fmt"
	"strconv"
)

func Calc(args []string) {
	if len(args) < 3 {
		fmt.Println("Error: insufficient arguments")
		fmt.Println("Usage: gocli calc <add|sub|mul|div|mod> <a> <b>")
		return
	}

	op, ok := ops[args[0]]
	if !ok {
		fmt.Printf("Error: unknown operation '%s'\n", args[0])
		fmt.Println("Supported operations: add, sub, mul, div, mod")
		return
	}

	a, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Printf("Error: invalid number '%s'\n", args[1])
		return
	}

	b, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Printf("Error: invalid number '%s'\n", args[2])
		return
	}

	if (args[0] == "div" || args[0] == "mod") && b == 0 {
		fmt.Println("Error: division by zero")
		return
	}

	result := op(a, b)
	fmt.Printf("%s %s %s = %v\n", args[1], args[0], args[2], result)
}

