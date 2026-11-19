package cmd

import (
	"fmt"
	"path/filepath"
	"servman/sv"

	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:     "disable",
	Aliases: []string{"d"},
	Short:   "Disable a service",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		disable(&args[0])
	},
}

func disable(name *string) {
	service := mustFind(*name)
	if !service.Active {
		fmt.Printf("%s is already disabled.\n", service.Name)
		return
	}
	// The entry in ActiveDir is a symlink, so plain rm (no -r) removes only the link.
	sudo("rm", filepath.Join(sv.ActiveDir, service.Name))
	fmt.Printf("%s disabled.\n", service.Name)
}

func init() {
	rootCmd.AddCommand(disableCmd)
}
