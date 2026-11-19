package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Mark a service as up on boot",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		up(&args[0])
	},
}

func up(name *string) {
	service := mustFind(*name)
	sudo("rm", "-f", filepath.Join(service.Location, service.Name, "down"))
	fmt.Printf("%s is up on boot.\n", service.Name)
}

func init() {
	rootCmd.AddCommand(upCmd)
}
