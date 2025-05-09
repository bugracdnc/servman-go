package cmd

import (
	"log"
	"servman/sv"

	"github.com/spf13/cobra"
)

var (
	enabled, disabled           bool = true, true
	_enabledFlag, _disabledFlag bool = false, false
)

var services = sv.FetchServices(sv.DefaultLocation)

var rootCmd = &cobra.Command{
	Use:   "servman [command] service",
	Short: "Enable or disable runit services with ease",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
