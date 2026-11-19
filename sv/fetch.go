package sv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getServices(location, activeDir string) ([]Service, error) {
	dirs, err := os.ReadDir(location)
	if err != nil {
		return nil, err
	}

	var services []Service
	for _, dir := range dirs {
		_, downErr := os.Stat(filepath.Join(location, dir.Name(), "down"))
		_, activeErr := os.Stat(filepath.Join(activeDir, dir.Name()))
		services = append(services, Service{
			Name:        dir.Name(),
			Active:      activeErr == nil,
			Location:    location,
			StartOnBoot: downErr != nil,
		})
	}
	return services, nil
}

// FetchServices lists the services installed in location.
func FetchServices(location string) ([]Service, error) {
	services, err := getServices(location, ActiveDir)
	if err != nil {
		return nil, fmt.Errorf("reading runit services (is this a runit system?): %w", err)
	}
	return services, nil
}

// Find returns the service whose name matches case-insensitively.
func Find(services []Service, name string) (Service, error) {
	for _, s := range services {
		if strings.EqualFold(s.Name, name) {
			return s, nil
		}
	}
	return Service{}, fmt.Errorf("no service named %q in %s", name, DefaultLocation)
}
