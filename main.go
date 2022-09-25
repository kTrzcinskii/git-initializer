package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func CheckErrors (err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitFlags() (dir, name *string, priv, github, readme *bool) {
	dirPtr := flag.String("dir", "/", "Directory in which you want to initialize your git repository")
	privatePtr := flag.Bool("priv", false, "Do you want this repository to be private?")
	namePtr := flag.String("name", "project", "Name of your project")
	createGithubRepoPtr := flag.Bool("github", false, "Do you want to create this repository on your github accout?")
	addReadmePtr := flag.Bool("readme", false, "Do you want to auto-create README.md file for this project?")

	flag.Parse()

	return dirPtr, namePtr, privatePtr, createGithubRepoPtr, addReadmePtr
}

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

func InitGit(path string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	fmt.Println(string(out))
	return nil
}

func main() {
	dir, name, priv, github, readme := InitFlags()

	path := fmt.Sprintf("%s/%s", *dir, *name)

	if CheckIfDirectoryAlreadyExists(path) {
		log.Printf("There is already directory \"%s\" in %s", *name, *dir)
		os.Exit(1)
	}

	err := CreateProjectDirectory(*dir, *name, path)
	CheckErrors(err)

	err = InitGit(path)
	CheckErrors(err)

	fmt.Println("dir: ", *dir)
	fmt.Println("private: ", *priv)
	fmt.Println("name: ", *name)
	fmt.Println("create github: ", *github)
	fmt.Println("add readme: ", *readme)
}