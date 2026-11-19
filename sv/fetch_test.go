package sv

import (
	"os"
	"path/filepath"
	"testing"
)

// Builds a fake runit layout: sshd enabled, cron enabled but down on boot, nginx installed only.
func fakeRunit(t *testing.T) (svDir, activeDir string) {
	t.Helper()
	root := t.TempDir()
	svDir, activeDir = filepath.Join(root, "etc/sv"), filepath.Join(root, "var/service")
	for _, name := range []string{"sshd", "cron", "nginx"} {
		if err := os.MkdirAll(filepath.Join(svDir, name), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	os.MkdirAll(activeDir, 0o755)
	os.WriteFile(filepath.Join(svDir, "cron", "down"), nil, 0o644)
	for _, name := range []string{"sshd", "cron"} {
		if err := os.Symlink(filepath.Join(svDir, name), filepath.Join(activeDir, name)); err != nil {
			t.Fatal(err)
		}
	}
	return svDir, activeDir
}

func TestGetServices(t *testing.T) {
	svDir, activeDir := fakeRunit(t)
	services, err := getServices(svDir, activeDir)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string][2]bool{ // name -> {Active, StartOnBoot}
		"sshd":  {true, true},
		"cron":  {true, false},
		"nginx": {false, true},
	}
	if len(services) != len(want) {
		t.Fatalf("got %d services, want %d", len(services), len(want))
	}
	for _, s := range services {
		if got := [2]bool{s.Active, s.StartOnBoot}; got != want[s.Name] {
			t.Errorf("%s: {Active, StartOnBoot} = %v, want %v", s.Name, got, want[s.Name])
		}
		if s.Location != svDir {
			t.Errorf("%s: Location = %q", s.Name, s.Location)
		}
	}
}

func TestFetchServicesMissingDir(t *testing.T) {
	if _, err := FetchServices(filepath.Join(t.TempDir(), "nope")); err == nil {
		t.Fatal("expected an error for a missing service directory")
	}
}

func TestFind(t *testing.T) {
	services := []Service{{Name: "sshd"}, {Name: "NetworkManager"}}
	if s, err := Find(services, "networkmanager"); err != nil || s.Name != "NetworkManager" {
		t.Errorf("Find case-insensitive = %v, %v", s, err)
	}
	if _, err := Find(services, "ssh"); err == nil {
		t.Error("Find should not match prefixes")
	}
}
