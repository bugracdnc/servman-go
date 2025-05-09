package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

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
	fmt.Printf("Disabling: '%s'\n", *name)
	for _, service := range *services {
		if strings.EqualFold(service.Name, *name) {
			cmd := exec.Command("sudo", "rm", "-rf", "/var/service/"+service.Name)
			stdout, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s disabled. %s\n", service.Name, stdout)
		}
	}
}

func init() {
	rootCmd.AddCommand(disableCmd)
}
