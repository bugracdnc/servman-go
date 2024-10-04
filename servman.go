package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Flags struct {
	//Search services
	search *string

	//Delimiters to list
	active   *bool
	disabled *bool

	//Enable or disable services by name
	enable  *string
	disable *string

	//up and down a service by name
	up   *string
	down *string
}

func handleFlags() *Flags {
	flags := Flags{}

	flags.search = flag.String("s", "", "Search flag (shorthand)")
	flag.StringVar(flags.search, "search", "", "Search flag")

	flags.active = flag.Bool("a", false, "Active flag")

	flags.disabled = flag.Bool("d", false, "Disabled flag")

	flag.StringVar(flags.search, "enable", "", "Search flag")

	flags.disable = flag.String("disable", "", "Enable flag")

	flags.down = flag.String("down", "", "Down flag")

	flags.up = flag.String("up", "", "Up flag")

	flag.Parse()

	if !*flags.disabled && !*flags.active {
		*flags.active = true
		*flags.disabled = true
	}

	return &flags
}

func main() {
	flags := handleFlags()
	defaultLocation := "/etc/sv"
	services := fetchServices(defaultLocation)

	if len(*flags.enable) > 0 {
		enable(services, flags.enable)
	} else if len(*flags.disable) > 0 {
		disable(services, flags.disable)
	} else if len(*flags.search) > 0 {
		search(services, flags.search, flags.active, flags.disabled)
	} else if len(*flags.up) > 0 {
		up(services, flags.up)
	} else if len(*flags.down) > 0 {
		down(services, flags.down)
	} else {
		list(services, flags.active, flags.disabled)
	}
}

func search(services *[]Service, name *string, activeFlag *bool, disabledFlag *bool) {
	fmt.Printf("Searching services: '%s'\n", *name)
	var foundServices []Service
	for _, service := range *services {
		if strings.Contains(service.Name, *name) {
			foundServices = append(foundServices, service)
		}
	}
	printServices(&foundServices, activeFlag, disabledFlag)
}

func list(services *[]Service, activeFlag *bool, disabledFlag *bool) {
	fmt.Print("Listing services: ")
	if *activeFlag {
		fmt.Print("active")
	}
	if *activeFlag && *disabledFlag {
		fmt.Print(", ")
	}
	if *disabledFlag {
		fmt.Print("disabled")
	}
	fmt.Println()
	printServices(services, activeFlag, disabledFlag)
}

func disable(services *[]Service, name *string) {
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

func enable(services *[]Service, name *string) {
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

func down(services *[]Service, name *string) {
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

func up(services *[]Service, name *string) {
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
