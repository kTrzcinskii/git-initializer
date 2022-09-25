package main

import "flag"

func InitFlags() (dir, name *string, priv, github, readme *bool) {
	dirPtr := flag.String("dir", "/", "Directory in which you want to initialize your git repository")
	privatePtr := flag.Bool("priv", false, "Do you want this repository to be private?")
	namePtr := flag.String("name", "project", "Name of your project")
	createGithubRepoPtr := flag.Bool("github", false, "Do you want to create this repository on your github accout?")
	addReadmePtr := flag.Bool("readme", false, "Do you want to auto-create README.md file for this project?")

	flag.Parse()

	return dirPtr, namePtr, privatePtr, createGithubRepoPtr, addReadmePtr
}