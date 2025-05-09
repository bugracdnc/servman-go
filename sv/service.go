package sv

var DefaultLocation = "/etc/sv"

type Service struct {
	Name        string
	Active      bool
	Location    string
	StartOnBoot bool
}
