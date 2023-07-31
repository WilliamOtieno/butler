package services

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

func GetOS() (string, error) {
	goos := runtime.GOOS
	log.Println("OS Detected: ", goos)
	switch goos {
	case "linux":
		log.Println("Attempting to get specific distro...")
		distro, err := getLinuxDistro()
		if err != nil {
			log.Println(err)
		}
		log.Println("Distro Detected: ", distro)
		return distro, nil
	default:
		return goos, nil
	}
}

func getLinuxDistro() (string, error) {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "ID=") {
			return strings.TrimPrefix(line, "ID="), nil
		}
	}

	return "", fmt.Errorf("distro ID not found in /etc/os-release")
}
