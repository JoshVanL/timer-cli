package main

import (
	"fmt"
	"os"

	"github.com/JoshVanL/timer-cli/pkg/timer"
)

func main() {
	args := os.Args[1:]

	t := timer.New()
	if err := t.ParseArguments(args); err != nil {
		fmt.Printf("Error parsing input: %v", err)
		os.Exit(1)
	}

	t.StartTimer()
	fmt.Printf("\nTimer over: %s\n", t.GetTimes())
}
