package services

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func findBaseDir(dir string) (string, error) {
	if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
		return dir, nil
	}

	parentDir := filepath.Dir(dir)
	if parentDir == dir {
		return "", os.ErrNotExist
	}

	return findBaseDir(parentDir)
}

func parseScriptPath(path string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	dir, err := findBaseDir(wd)
	if err != nil {
		log.Println(err)
	}
	finalPath := dir + path
	return finalPath
}

func runScript(scriptPath string) {
	cmd := exec.Command("sh", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
