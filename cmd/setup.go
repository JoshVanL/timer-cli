package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/JoshVanL/timer-cli/pkg/timer"
)

var RootCmd = &cobra.Command{
	Use:   "timer [duration]",
	Short: "Start a clock timer. Inputs are number duration suffixed by a time value\ns-second m-minute h-hour. e.g. timer 4s 17m 1h\n Order does not matter. Duration with no suffix will be used as seconds",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("No arguments given.")
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
