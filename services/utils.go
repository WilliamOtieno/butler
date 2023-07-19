package services

import (
	"log"
	"os"
	"os/exec"
)

func runScript(commands []string) {
	for _, cmdStr := range commands {
		cmd := exec.Command("sh", "-c", cmdStr)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
	}
}
