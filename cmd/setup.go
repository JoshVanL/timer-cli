package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/JoshVanL/timer-cli/pkg/timer"
)

var short = "Start a cli clock timer. Inputs are number duration suffixed by a time value\ns-second m-minute h-hour. e.g. timer 4s 17m 1h\nOrder does not matter. Duration with no suffix will be used as seconds"
var use = "timer [duration]"

var RootCmd = &cobra.Command{
	Use:   use,
	Short: short,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Printf("No arguments given.\n\n%s\n\nUsage:\n	%s\n", short, use)
			os.Exit(1)
		}

		if len(args) > 3 {
			fmt.Println("Too many arguments given. (>3)")
			os.Exit(1)
		}

		t := timer.New()
		if err := t.ParseArguments(args); err != nil {
			fmt.Printf("Error parsing input: %v", err)
			os.Exit(1)
		}

		t.StartTimer()
		fmt.Printf("\nTimer over: %s\n", t.GetTimes())
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
