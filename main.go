package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CheckErrors (err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dir, name, priv, createGithub, readme := InitFlags()

	path := filepath.Join(*dir, *name)

	if CheckIfDirectoryAlreadyExists(path) {
		log.Printf("There is already directory \"%s\" in %s", *name, *dir)
		os.Exit(1)
	}

	githubAccessToken, err := LoadGithubAccessToken()
	CheckErrors(err)

	if *createGithub && githubAccessToken == "" {
		log.Fatal("Enter your github access token in .env file to create github repository!")
		os.Exit(1)
	}

	err = CreateProjectDirectory(*dir, *name, path)
	CheckErrors(err)

	err = InitGit(path)
	CheckErrors(err)

	if *readme {
		err = AddReadme(path, *name)
		CheckErrors(err)
	}

	if *createGithub {
		err = CreateAndConnectGithub(*name, path, *priv)
		CheckErrors(err)
	}

	fmt.Println("dir: ", *dir)
	fmt.Println("private: ", *priv)
	fmt.Println("name: ", *name)
	fmt.Println("create github: ", *createGithub)
	fmt.Println("add readme: ", *readme)
}