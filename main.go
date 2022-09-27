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
	flags, err := InitFlags()
	CheckErrors(err)

	path := filepath.Join(flags.dir, flags.name)

	if CheckIfDirectoryAlreadyExists(path) {
		log.Printf("There is already directory \"%s\" in %s", flags.name, flags.dir)
		os.Exit(1)
	}

	githubAccessToken, err := LoadGithubAccessToken()
	CheckErrors(err)

	if flags.github && githubAccessToken == "" {
		log.Fatal("Enter your github access token in .env file to create github repository!")
		os.Exit(1)
	}

	err = CreateProjectDirectory(flags.dir, flags.name, path)
	CheckErrors(err)

	err = InitGit(path)
	CheckErrors(err)

	if flags.readme {
		err = AddReadme(path, flags.name)
		CheckErrors(err)
	}

	if flags.github {
		err = CreateAndConnectGithub(flags.name, path, githubAccessToken, flags.priv)
		CheckErrors(err)
	}

	fmt.Println("Everything is ready. Happy coding!")
}