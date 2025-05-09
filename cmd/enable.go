package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

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
	fmt.Printf("Enabling: '%s'\n", *name)
	for _, service := range *services {
		if strings.EqualFold(service.Name, *name) {
			cmd := exec.Command("sudo", "ln", "-s", "/etc/sv/"+service.Name, "/var/service")
			stdout, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s enabled. %s\n", service.Name, stdout)
		}
	}
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
