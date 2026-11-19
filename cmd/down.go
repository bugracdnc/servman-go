package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Mark a service as down on boot",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		down(&args[0])
	},
}

func down(name *string) {
	service := mustFind(*name)
	sudo("touch", filepath.Join(service.Location, service.Name, "down"))
	fmt.Printf("%s is down on boot.\n", service.Name)
}

func init() {
	rootCmd.AddCommand(downCmd)
}
