package cmd

import (
	"fmt"
	"strings"
)

func Greet(args []string) {
	if len(args) == 0 {
		fmt.Println("Hello, World! :)")
		return
	}

	name := args[0]
	greeting := "Hello"

	if len(args) > 1 {
		switch strings.ToLower(args[1]) {
		case "formal":
			greeting = "Greetings"
		case "casual":
			greeting = "Hey"
		case "friendly":
			greeting = "Hi"
		}
	}

	fmt.Printf("%s, %s! :)\n", greeting, name)
}