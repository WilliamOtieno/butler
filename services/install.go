package services

import (
	"github.com/williamotieno/butler/scripts/install/docker"
	"github.com/williamotieno/butler/scripts/install/nvm"
	"github.com/williamotieno/butler/scripts/install/redis"
	"log"
)

func InstallTest() {
	runScript(docker.GetTestCommands())
}

func InstallDocker() {
	log.Println("Installing Docker: ...")
	os, err := GetOS()
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
	default:
		log.Println("Error: Not yet implemented on platform")
	}
}

func InstallRedis() {
	log.Println("Installing Docker: ...")
	os, err := GetOS()
	if err != nil {
		log.Println(err)
	}

	switch os {
	case "debian":
		runScript(redis.GetDebianCommands())
	case "ubuntu":
		runScript(redis.GetUbuntuCommands())
	case "darwin":
		runScript(redis.GetDarwinCommands())
	case "windows":
		log.Println("Error: Redis Installation on windows requires WSL. Use debian/ubuntu installation ")
	default:
		log.Println("Error: Not yet implemented on platform")
	}
}

func InstallNVM() {
	runScript(nvm.GetNVMCommands())
}
