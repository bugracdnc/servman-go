package cmd

import (
	"log"
	"os/exec"
	"servman/sv"

	"github.com/spf13/cobra"
)

var (
	enabled, disabled           bool = true, true
	_enabledFlag, _disabledFlag bool = false, false
)

// Loaded on demand so --help and completion work on systems without runit.
func services() []sv.Service {
	s, err := sv.FetchServices(sv.DefaultLocation)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func mustFind(name string) sv.Service {
	s, err := sv.Find(services(), name)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func sudo(args ...string) {
	out, err := exec.Command("sudo", args...).CombinedOutput()
	if err != nil {
		log.Fatalf("%v: %s", err, out)
	}
}

var rootCmd = &cobra.Command{
	Use:   "servman [command] service",
	Short: "Enable or disable runit services with ease",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
