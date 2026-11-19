package cmd

import (
	"fmt"
	"path/filepath"
	"servman/sv"

	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:     "enable",
	Aliases: []string{"en"},
	Short:   "Enable a service",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		enable(&args[0])
	},
}

func enable(name *string) {
	service := mustFind(*name)
	if service.Active {
		fmt.Printf("%s is already enabled.\n", service.Name)
		return
	}
	sudo("ln", "-s", filepath.Join(service.Location, service.Name), sv.ActiveDir)
	fmt.Printf("%s enabled.\n", service.Name)
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
