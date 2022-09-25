package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)


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

func LoadGithubAccessToken() (string, error) {
	err := godotenv.Load()

	if err != nil {
		return "", err
	}

	at := os.Getenv("GITHUB_ACCESS_TOKEN")

	return at, nil
}

//https://github.com/google/go-github
func CreateAndConnectGithub(name, path string, priv bool) error {
	panic("todo")
}