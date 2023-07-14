package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
)

func getOS() (string, error) {
	goos := runtime.GOOS
	log.Println("OS Detected: ", goos)
	if goos == "linux" {
		log.Print("Attempting to get specific distro...")
		distro, err := getLinuxDistro()
		if err != nil {
			log.Println(err)
		}
		log.Println("Distro Detected: ", distro)
		return distro, nil
	}
	return goos, nil
}

func getLinuxDistro() (string, error) {
	data, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "ID=") {
			return strings.TrimPrefix(line, "ID="), nil
		}
	}

	return "", fmt.Errorf("Distro ID not found in /etc/os-release")
}
