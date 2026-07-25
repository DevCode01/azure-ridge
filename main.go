package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("$NAME v0.1.0 — $DESC")
		fmt.Println("Usage: $NAME [options]")
		os.Exit(0)
	}
	fmt.Printf("Processing: %v\n", args)
	// TODO: implement core logic
}
