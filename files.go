package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

	fmt.Println()
	fmt.Println("Project directory has been created at", path)
	fmt.Println()
	return nil
}

func AddReadme(path, name string) error {
	fileName := filepath.Join(path, "README.md")
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	_, err = file.Write([]byte("# Repository created with Git initializer.\nYou can find more about it [here](https://github.com/kTrzcinskii/git-initializer)."))

	if err != nil {
		return err
	}

	err = file.Close()

	if err != nil {
		return err
	}

	fmt.Println("README.md file has been created. Remember to change it!")

	return nil
}