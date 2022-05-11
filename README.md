<img src="https://user-images.githubusercontent.com/64161383/155763268-e09d9613-a53f-4ec7-a943-aab93ef2ffa6.png" width="150px" alt="logo"  align="right" />

<div align="left">

 <h1> borzoi </h1>
 
 <a href="https://github.com/automation-co/borzoi/actions?query=branch%3Amain"><img src="https://github.com/automation-co/borzoi/workflows/Go/badge.svg?branch=main" /> </a> <a href="https://github.com/automation-co/borzoi/releases"> <img src="https://img.shields.io/github/release/automation-co/borzoi.svg" /> </a> <img src="https://img.shields.io/github/go-mod/go-version/automation-co/borzoi" /> <a href="https://goreportcard.com/report/github.com/automation-co/borzoi"><img src="https://goreportcard.com/badge/github.com/automation-co/borzoi" /> </a> <img src="https://img.shields.io/github/license/automation-co/borzoi" /> <img src="https://img.shields.io/github/issues/automation-co/borzoi" />

 </div>

<!-- --- -->

**Make polyrepos easy!**

Polyrepos are as cool as monorepos, if managed well.

## Docs

### Installation

```
go install github.com/automation-co/borzoi@latest
```

### Getting Started

Borzoi helps you to easily replicate and share file structure.

You can use the `borzoi` command to run the borzoi command line tool.

```
Borzoi is a tool that makes it easy to manage your codebase.

It helps you a simple interface to standardize your git repos in the same manner.

For more information, please visit
https://github.com/automation-co/borzoi

Usage:
  borzoi [command]

Available Commands:
  clone       Clones the repos
  completion  Generate the autocompletion script for the specified shell
  freeze      Generates borzoi-lock.json
  generate    Generates the config file
  help        Help about any command

Flags:
  -h, --help     help for borzoi
  -t, --toggle   Help message for toggle

Use "borzoi [command] --help" for more information about a command.
```

### Authentication

If you want to use borzoi with a private repository, you can use the `--access` flag to authenticate.
Pass the access token as the value of the flag.

> Make sure the personal access token has repo scope.

If you need help, creating the personal access token, please visit the [GitHub API documentation](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)

If you don't have a global git username or you want to use a different username, you can use the `--username` flag to authenticate.
Pass the username as the value of the flag.

```
borzoi clone --access <access_token> --username <username>
```

### Ignore

If you want to exclude certain directories when using the `generate` command, you can add a `.borzoiignore` file which can have a list of subdirectories whose path you wish to exclude.

---



### Example

```javascript
// borzoi.json
{
  "automation-co/borzoi": "https://github.com/automation-co/borzoi.git",
  "automation-co/husky": "https://github.com/automation-co/husky.git",
  "dependencies/air": "https://github.com/cosmtrek/air.git",
  "dependencies/go-git": "https://github.com/go-git/go-git.git"
}

```

`borzoi clone` using this config file will result in

```
.
├── automation-co
│   ├── borzoi
│   └── husky
├── borzoi.json
└── dependencies
    ├── air
    └── go-git
```

---

<div align="center">

 Developed by <a href="https://github.com/automation-co" >@automation-co</a>

</div>
