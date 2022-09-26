package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Flags struct {
	dir string
	name string
	priv bool
	readme bool
	github bool
}

func LoadDefaultOptions () (defaults Flags, e error) {
	err := godotenv.Load()
	if err != nil {
		return Flags{}, err
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
		return Flags{}, err
		}
	}

	defaultReadmeString := os.Getenv("DEFAULT_README")
	var defaultReadme bool
	if (defaultReadmeString == "") {
		defaultReadme = false
	} else {
		defaultReadme, err = strconv.ParseBool(defaultReadmeString)
		if err != nil {
		return Flags{}, err
		}
	}

	defaultGithubString := os.Getenv("DEFAULT_GITHUB")
	var defaultGithub bool
	if (defaultGithubString == "") {
		defaultGithub = false
	} else {
		defaultGithub, err = strconv.ParseBool(defaultGithubString)
		if err != nil {
		return Flags{}, err
		}
	}

	return Flags{dir: defaultDir, name: defaultName, priv: defaultPriv, readme: defaultReadme, github: defaultGithub}, nil
}

func InitFlags() (f Flags, e error) {
	defaults, err := LoadDefaultOptions()
	if err != nil {
		return Flags{}, err
	}

	dirPtr := flag.String("dir", defaults.dir, "Directory in which you want to initialize your git repository")
	privatePtr := flag.Bool("priv", defaults.priv, "Do you want this repository to be private?")
	namePtr := flag.String("name", defaults.name, "Name of your project")
	createGithubRepoPtr := flag.Bool("github", defaults.github, "Do you want to create this repository on your github accout?")
	addReadmePtr := flag.Bool("readme", defaults.readme, "Do you want to auto-create README.md file for this project?")

	flag.Parse()

	return Flags{dir: *dirPtr, name: *namePtr, priv: *privatePtr, readme: *addReadmePtr, github: *createGithubRepoPtr}, nil
}