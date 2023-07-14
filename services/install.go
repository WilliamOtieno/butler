package services

import (
	"log"
)

func InstallDocker() {
	log.Println("Installing Docker: ...")
	dockerScripts := parseScriptPath("/scripts/install/docker/")
	os, err := getOS()
	if err != nil {
		log.Println(err)
	}
	switch os {
	case "debian":
		runScript(dockerScripts + "debian.sh")
	case "centos":
		runScript(dockerScripts + "centos.sh")
	case "fedora":
		runScript(dockerScripts + "fedora.sh")
	case "ubuntu":
		runScript(dockerScripts + "ubuntu.sh")
	case "darwin":
		log.Println("Error: Docker installation in MacOS is through docker-desktop")
	case "windows":
		log.Println("Error: Windows installation not yet implemented")
	}
}

func InstallTest() {
	dockerScripts := parseScriptPath("/scripts/install/docker/")
	runScript(dockerScripts + "test.sh")
}
