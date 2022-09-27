package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/google/go-github/v47/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
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

func CreateAndConnectGithub(name, path, accessToken string, priv bool) error {
	//login to github account
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	//create repo
	newRepo := &github.Repository{
			   Name: github.String(name),
			   Private: github.Bool(priv),
	}
	repo, _, err := client.Repositories.Create(ctx, "", newRepo)
	if err != nil {
		return err
	}

	gitUrl := repo.GetGitURL()
	httpsUrl := strings.Replace(gitUrl, "git", "https", 1)
	fmt.Println()
	fmt.Println("Your github repository has been created at:", httpsUrl)

	//push code to repo
	cmd := exec.Command("git", "remote", "add", "origin", httpsUrl)
	cmd.Dir = path
	_, err = cmd.CombinedOutput()

	if err != nil {
		return err
	}

	cmd = exec.Command("git", "branch", "-M", "main")
	cmd.Dir = path
	_, err = cmd.CombinedOutput()

	if err != nil {
		return err
	}

	cmd = exec.Command("git", "add", ".")
	cmd.Dir = path
	_, err = cmd.CombinedOutput()

	if err != nil {
		return err
	}

	cmd = exec.Command("git", "commit", "-m", "INIT")
	cmd.Dir = path
	_, err = cmd.CombinedOutput()

	if err != nil {
		return err
	}

	cmd = exec.Command("git", "push", "origin", "main")
	cmd.Dir = path
	_, err = cmd.CombinedOutput()
	
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("The initial commit has been pushed to your repository.")
	return nil
}