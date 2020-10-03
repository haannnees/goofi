<img src="https://github.com/haannnees/goofi/blob/master/doc/logo.png" alt="Goofi" width="163" height="274">

## Go **+** Find = **Goofi**❣️

A simple CLI-Tool to look up abbreviations built with love in [Go](https://golang.org/) with [Cobra](https://github.com/spf13/cobra).

## Overview
Goofi is a open source CLI-Tool written in GO, which can be used to find easy and quick the meaning of abbreviations.

## Background
In every company there are thousands of abbreviations that only long-time employees know. 
That's why I developed this library, because I wanted to create a solution that would allow new employees to quickly look up specific abbreviations.
In my current company there is a inner source Github Repository with a glossary, so i thought a CLI-Tool would be handy to access these quickly via CLI.

#### Prerequisite Tools
* [Git](https://git-scm.com/)
* [Go](https://golang.org/dl/)

#### Clone repository from Github
```bash
mkdir $HOME/workspace
cd $HOME/workspace
git clone https://github.com/haannnees/goofi.git
cd goofi
go install goofi
```

#### Run Goofi
```bash
goofi --help
```

#### Import acronyms
```bash
goofi add -s CLI=Command Line Interface
goofi add -r https://raw.githubusercontent.com/haannnees/goofi/master/doc/goofi.csv
goofi add -l $HOME/abbreviationList.csv
```

#### Find word
Enter command
```bash
goofi find CLI
```
Get result printed
```bash
┌─────────────────────────┐
│ RESULT(S) FOR CLI:      │
├─────────────────────────┤
│  Command Line Interface │
└─────────────────────────┘
```

#### Made by
Hannes Wagner             
Software Engineer         
[Github](https://github.com/haannnees)              
[LinkedIn](https://www.linkedin.com/in/hannes-wagner-171549128/)          


[![GitHub Super-Linter](https://github.com/haannnees/goofi/workflows/Lint%20Code%20Base/badge.svg)](https://github.com/marketplace/actions/super-linter)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/haannnees/goofi)