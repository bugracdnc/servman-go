package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

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
	fmt.Printf("Turning down: '%s'\n", *name)
	for _, service := range *services {
		if strings.EqualFold(service.Name, *name) {
			cmd := exec.Command("sudo", "touch", "/var/service/"+service.Name+"/down")
			stdout, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s is down on boot. %s\n", service.Name, stdout)
		}
	}
}

func init() {
	rootCmd.AddCommand(downCmd)
}
