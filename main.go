package main

import (
	"laba-3/utils"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: main [consumer|producer]")
		return
	}

	command := os.Args[1]
	if command == "consumer" {
		utils.Consumer()
	} else if command == "producer" {
		utils.Producer()
	} else {
		println("Unknown command")
	}
}
