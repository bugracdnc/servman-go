package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

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
	fmt.Printf("Turning up: '%s'\n", *name)
	for _, service := range *services {
		if strings.EqualFold(service.Name, *name) {
			cmd := exec.Command("sudo", "rm", "-f", "/var/service/"+service.Name+"/down")
			stdout, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s is up on boot. %s\n", service.Name, stdout)
		}
	}
}

func init() {
	rootCmd.AddCommand(upCmd)
}
