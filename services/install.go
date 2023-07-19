package services

import (
	"github.com/williamotieno/butler/scripts/install/docker"
	"log"
)

func InstallDocker() {
	log.Println("Installing Docker: ...")
	os, err := getOS()
	if err != nil {
		log.Println(err)
	}
	switch os {
	case "debian":
		runScript(docker.GetDebianCommands())
	case "centos":
		runScript(docker.GetCentosCommands())
	case "fedora":
		runScript(docker.GetFedoraCommands())
	case "ubuntu":
		runScript(docker.GetUbuntuCommands())
	case "darwin":
		log.Println("Error: Docker installation in MacOS is through docker-desktop")
	case "windows":
		log.Println("Error: Windows installation not yet implemented")
	}
}

func InstallTest() {
	runScript(docker.GetTestCommands())
}
