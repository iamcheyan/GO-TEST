package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const version = "v2.0.0"

var ops = map[string]func(float64, float64) float64{
	"add": func(a, b float64) float64 { return a + b },
	"sub": func(a, b float64) float64 { return a - b },
	"mul": func(a, b float64) float64 { return a * b },
	"div": func(a, b float64) float64 { return a / b },
	"mod": func(a, b float64) float64 { return float64(int(a) % int(b)) },
}

func mustParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("invalid number: %s", s)
	}
	return f
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Hello, World!")
		return
	}

	switch cmd, args := os.Args[1], os.Args[2:]; cmd {
	case "greet":
		name := "World"
		if len(args) > 0 {
			name = args[0]
		}
		fmt.Printf("Hello, %s!\n", name)

	case "calc":
		if len(args) < 3 {
			log.Fatal("usage: gocli calc <add|sub|mul|div|mod> <a> <b>")
		}
		op, ok := ops[args[0]]
		if !ok {
			log.Fatalf("unknown op: %s", args[0])
		}
		a, b := mustParseFloat(args[1]), mustParseFloat(args[2])
		if (args[0] == "div" || args[0] == "mod") && b == 0 {
			log.Fatal("division by zero")
		}
		fmt.Printf("%s %s %s = %v\n", args[1], args[0], args[2], op(a, b))

	case "upper":
		if len(args) < 1 {
			log.Fatal("usage: gocli upper <string>")
		}
		fmt.Println(strings.ToUpper(args[0]))

	case "version":
		fmt.Println("gocli " + version)

	case "help":
		fmt.Print(`gocli - a simple Go CLI

Usage: gocli <command> [arguments]

Commands:
  greet [name]                         Print a greeting
  calc <add|sub|mul|div|mod> <a> <b>   Arithmetic operations
  upper <string>                       Convert to uppercase
  version                              Print version
  help                                 Show this help
`)

	default:
		log.Fatalf("unknown command: %s\nRun 'gocli help' for usage.", cmd)
	}
}
