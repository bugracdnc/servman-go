package sv

import (
	"fmt"

	"github.com/fatih/color"
)

var green = color.New(color.FgGreen)
var blue = color.New(color.FgBlue)
var red = color.New(color.FgRed)

func printService(index int, color *color.Color, name string) int {
	fmt.Printf("%02d) ", index)
	n, err := color.Printf("%s\n", name)
	if err != nil {
		return -1
	}
	return n
}

func PrintServices(services *[]Service, activeFlag *bool, disabledFlag *bool) {
	index := 0
	for _, service := range *services {
		if !service.Active && *disabledFlag {
			index++
			printService(index, red, service.Name)
		} else if !service.StartOnBoot {
			index++
			printService(index, blue, service.Name)
		} else if service.Active && *activeFlag {
			index++
			printService(index, green, service.Name)
		}
	}
}
