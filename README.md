<img src="https://user-images.githubusercontent.com/64161383/155763268-e09d9613-a53f-4ec7-a943-aab93ef2ffa6.png" width="150px" alt="logo"  align="right" />

<div align="left">

# borzoi

[![Build Status](https://github.com/automation-co/borzoi/workflows/Go/badge.svg?branch=main)](https://github.com/automation-co/borzoi/actions?query=branch%3Amain)
[![Release](https://img.shields.io/github/release/automation-co/borzoi.svg)](https://github.com/automation-co/borzoi/releases)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/automation-co/borzoi)
[![Go Report Card](https://goreportcard.com/badge/github.com/automation-co/borzoi)](https://goreportcard.com/report/github.com/automation-co/borzoi)
![GitHub](https://img.shields.io/github/license/automation-co/borzoi)
![GitHub issues](https://img.shields.io/github/issues/automation-co/borzoi)

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
  generate    Generates the config file
  help        Help about any command

Flags:
  -h, --help     help for borzoi
  -t, --toggle   Help message for toggle

Use "borzoi [command] --help" for more information about a command.
```

---

### Example

```javascript
// borzoi.json
{
  "automation-co": {
    "borzoi": {
      "repo": "https://github.com/automation-co/borzoi.git"
    },
    "husky": {
      "repo": "https://github.com/automation-co/husky.git"
    }
  }
}

```

will result in

```
.
├── automation-co
│   ├── borzoi  
│   └── husky  
└── borzoi.json
```

---

<div align="center">

Developed by [@automation-co](https://github.com/automation-co)

</div>
