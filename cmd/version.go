package cmd

import "fmt"

func Version() {
	fmt.Println("gocli " + versionStr)
}