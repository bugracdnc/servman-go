package sv

// DefaultLocation holds every installed runit service; ActiveDir holds symlinks to the enabled ones.
var (
	DefaultLocation = "/etc/sv"
	ActiveDir       = "/var/service"
)

type Service struct {
	Name        string
	Active      bool
	Location    string
	StartOnBoot bool
}
