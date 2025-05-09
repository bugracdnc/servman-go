package cmd

import (
	"fmt"
	"servman/sv"
	"strings"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"se"},
	Short:   "Search for services",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		disabled = !_enabledFlag
		enabled = !_disabledFlag
		search(&args[0], &enabled, &disabled)
	},
}

func search(name *string, enabledFlag *bool, disabledFlag *bool) {
	fmt.Printf("Searching services: '%s'\n", *name)
	var foundServices []sv.Service
	for _, service := range *services {
		if strings.Contains(service.Name, *name) {
			foundServices = append(foundServices, service)
		}
	}
	sv.PrintServices(&foundServices, enabledFlag, disabledFlag)
}

func init() {
	searchCmd.Flags().BoolVarP(&_enabledFlag, "enabled", "e", false, "Search only enabled services")
	searchCmd.Flags().BoolVarP(&_disabledFlag, "disabled", "d", false, "Search only disabled services")
	rootCmd.AddCommand(searchCmd)
}
