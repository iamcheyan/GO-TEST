package main

import "strings"

func greet(name string) string {
	return "Hello, " + name
}

func splitWords(text string) []string {
	return strings.Split(text, " ")
}
