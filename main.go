package main

import "flag"

func handleFlags() (*bool, *bool, *bool) {
	activeFlag := flag.Bool("a", false, "Active flag (shorthand)")
	flag.BoolVar(activeFlag, "active", false, "Active flag")

	disabledFlag := flag.Bool("d", false, "Disabled flag  (shorthand)")
	flag.BoolVar(disabledFlag, "disabled", false, "Disabled flag")

	downFlag := flag.Bool("down", false, "Down flag")

	flag.Parse()

	if !*activeFlag && !*disabledFlag && !*downFlag {
		trueRef := true
		activeFlag = &trueRef
		disabledFlag = &trueRef
		downFlag = &trueRef
	}

	return activeFlag, disabledFlag, downFlag
}

func main() {
	activeFlag, disabledFlag, downFlag := handleFlags()

	defaultLocation := "/etc/sv"

	printServices(fetchServices(defaultLocation), activeFlag, disabledFlag, downFlag)
}
