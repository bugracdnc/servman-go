package sv

import (
	"log"
	"os"
)

func getServices(location string) (*[]Service, error) {
	var services []Service
	activeServDir := "/var/service"

	dirs, err := os.ReadDir(location)
	if err != nil {
		return nil, err
	}

	for _, dir := range dirs {
		startOnBoot := true
		if _, err := os.ReadFile(location + "/" + dir.Name() + "/" + "down"); err == nil {
			startOnBoot = false
		}
		isActive := false
		if _, err := os.ReadDir(activeServDir + "/" + dir.Name()); err == nil {
			isActive = true
		}
		services = append(services, Service{Name: dir.Name(), Active: isActive, Location: location, StartOnBoot: startOnBoot})

	}
	return &services, nil
}

func FetchServices(location string) *[]Service {

	services, err := getServices(location)
	if err != nil {
		log.Fatal(err)
	}

	return services

}
