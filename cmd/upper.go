package cmd

import (
	"fmt"
	"log"
	"strings"
)

func Upper(args []string) {
	if len(args) < 1 {
		log.Fatal("usage: gocli upper <string>")
	}
	fmt.Println(strings.ToUpper(args[0]))
}