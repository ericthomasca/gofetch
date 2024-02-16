package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// info "Host" model
	// info "Kernel" kernel
	// info "Uptime" uptime
	// info "Packages" packages
	// info "Shell" shell
	// info "Resolution" resolution
	// info "DE" de
	// info "WM" wm
	// info "WM Theme" wm_theme
	// info "Theme" theme
	// info "Icons" icons
	// info "Terminal" term
	// info "Terminal Font" term_font
	// info "CPU" cpu
	// info "GPU" gpu
	// info "Memory" memory
	title, err := getTitle()
	if err != nil {
		log.Fatal(err)
	}

	distro, err := getDistro()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
	fmt.Println()
	fmt.Printf("OS: %s\n", distro)

}

func getTitle() (string, error) {
	userBytes, err := exec.Command("id", "-un").Output()
	if err != nil {
		return "", fmt.Errorf("error getting username: %v", err)
	}
	user := strings.TrimSpace(string(userBytes))

	hostnameBytes, err := exec.Command("hostname", "-f").Output()
	if err != nil {
		return "", fmt.Errorf("error getting hostname: %v", err)
	}
	hostname := strings.TrimSpace(string(hostnameBytes))

	title := fmt.Sprintf("%s@%s", user, hostname)
	if title == "" {
		return "", fmt.Errorf("problem getting title but no error generated")
	}

	return title, nil
}

func parseOsRelease(filename string) (map[string]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	result := make(map[string]string)
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(strings.Trim(parts[1], `"`))
			result[key] = value
		}
	}

	return result, nil
}

func getDistro() (string, error) {
	osReleaseData, err := parseOsRelease("/etc/os-release")
	if err != nil {
		return "", err
	}

	distro_name := osReleaseData["PRETTY_NAME"]

	return distro_name, nil
}
