package main

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckIfDirectoryAlreadyExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func CreateProjectDirectory(dir, name, path string) error {
	cmd := exec.Command("mkdir", name)
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	fmt.Println("Project directory has been created at", path)
	return nil
}