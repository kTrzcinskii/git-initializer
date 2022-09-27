# Git Initializer

## What is it?

It's a cli program written in Go that automates the process of creating new programming projects. It can:
 - create new project directory on your machine
 - add empty readme file
 - create the github repostiory with the same name and make the initial commit

I created it as a way to learn more about the golang. And this is the reason why I try to use as few libraries as possible (e.g. I didn't use [Cobra](https://cobra.dev/)).

## How to use it?

1. ```git clone https://github.com/kTrzcinskii/git-initializer```

2. ```cd git-initializer```

3. ```go build .```

4. ```./git-initializer <flags>```

e.g. ```./git-initializer -dir="/c/users/me/projects" -name="important_stuff" -priv=false -readme=true -github=true```

## Flags

All flags are optional, you can use git initializer with default options.

| Name | About | Input type | Default |
| ---- | ----- | ---------- | ------- |
| dir | It specifies in which directory on your machine you want to create the project | String, e.g. "/c/users/USERNAME/github/repos" | "/c" |
| name | It's the name of your project directory and the github repo (if you decide to create one) | String, e.g. "Project" | "project" |
| priv | It specifies whether the Github repo will be private or not. If github repo is not created, this flag doesn't change anything | Boolean | false |
| readme | It specifies if you want to create default README.md file (that you will have to edit yourself later) | Boolean | false |
| github | It specifies if you want to create github repository for this project (it requires you to enter github access token in .env file) | Boolean | false |

## How to change default options?

It's very easy, you just have to enter the default values that you want the program to have in .env file,
e.g.

```
DEFAULT_DIR="STRING"
DEFAULT_NAME="STRING"
DEFAULT_PRIV="BOOL(TRUE/FALSE)"
DEFAULT_GITHUB="BOOL(TRUE/FALSE)"
DEFAULT_README="BOOL(TRUE/FALSE)"
```

## Adding Github Access Token

If you want the git initializer to be able to create github repo for you, you must make sure that you provide it with your personal github access token.
You have to enter it to your .env file (nobody but you have access to it, so your access token is safe there)

```
GITHUB_ACCESS_TOKEN="STRING"
```
