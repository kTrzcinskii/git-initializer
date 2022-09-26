package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadDefaultOptions () (dir, name string, priv, readme, github bool, e error) {
	err := godotenv.Load()
	if err != nil {
		return "", "", false, false, false, err
	}
	defaultDir := os.Getenv("DEFAULT_DIR")
	if defaultDir == "" {
		defaultDir = "/c"
	}

	defaultName := os.Getenv("DEFAULT_NAME")
	if defaultName == "" {
		defaultName = "project"
	}

	defaultPrivString := os.Getenv("DEFAULT_PRIV")
	var defaultPriv bool
	if (defaultPrivString == "") {
		defaultPriv = false
	} else {
		defaultPriv, err = strconv.ParseBool(defaultPrivString)
		if err != nil {
		return "", "", false, false, false, err
	}
	}

	defaultReadmeString := os.Getenv("DEFAULT_README")
	var defaultReadme bool
	if (defaultReadmeString == "") {
		defaultReadme = false
	} else {
		defaultReadme, err = strconv.ParseBool(defaultReadmeString)
		if err != nil {
		return "", "", false, false, false, err
	}
	}

	defaultGithubString := os.Getenv("DEFAULT_GITHUB")
	var defaultGithub bool
	if (defaultGithubString == "") {
		defaultGithub = false
	} else {
		defaultGithub, err = strconv.ParseBool(defaultGithubString)
		if err != nil {
		return "", "", false, false, false, err
	}
	}

	return defaultDir, defaultName, defaultPriv, defaultReadme, defaultGithub, nil
}

func InitFlags() (dir, name *string, priv, github, readme *bool, e error) {
	defaultDir, defaultName, defaultPriv, defaultReadme, defaultGithub, err := LoadDefaultOptions()
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	dirPtr := flag.String("dir", defaultDir, "Directory in which you want to initialize your git repository")
	privatePtr := flag.Bool("priv", defaultPriv, "Do you want this repository to be private?")
	namePtr := flag.String("name", defaultName, "Name of your project")
	createGithubRepoPtr := flag.Bool("github", defaultGithub, "Do you want to create this repository on your github accout?")
	addReadmePtr := flag.Bool("readme", defaultReadme, "Do you want to auto-create README.md file for this project?")

	flag.Parse()

	return dirPtr, namePtr, privatePtr, createGithubRepoPtr, addReadmePtr, nil
}