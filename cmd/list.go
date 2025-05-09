package cmd

import (
	"fmt"
	"servman/sv"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List services",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		disabled = !_enabledFlag
		enabled = !_disabledFlag
		list(&enabled, &disabled)
	},
}

func list(enabledFlag *bool, disabledFlag *bool) {
	fmt.Print("Listing services: ")
	if *enabledFlag {
		fmt.Print("enabled")
	}
	if *enabledFlag && *disabledFlag {
		fmt.Print(", ")
	}
	if *disabledFlag {
		fmt.Print("disabled")
	}
	fmt.Println()
	sv.PrintServices(services, enabledFlag, disabledFlag)
}

func init() {
	listCmd.Flags().BoolVarP(&_enabledFlag, "enabled", "e", false, "List only enabled services")
	listCmd.Flags().BoolVarP(&_disabledFlag, "disabled", "d", false, "List only disabled services")
	rootCmd.AddCommand(listCmd)

}
