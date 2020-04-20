# confex

[![Go Report Card](https://goreportcard.com/badge/github.com/oncilla/confex)](https://goreportcard.com/report/github.com/oncilla/confex)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/oncilla/confex)
[![Release](https://img.shields.io/github/release-pre/oncilla/confex.svg)](https://github.com/oncilla/confex/releases)
[![license](https://img.shields.io/github/license/oncilla/confex.svg?maxAge=2592000)](https://github.com/oncilla/confex/blob/master/LICENSE)

confex helps you explore configuration files. It currently supports yaml, json,
and toml configurations.

## Installation

    go get github.com/oncilla/confex@latest


In case you see the following error:

    go: cannot use path@version syntax in GOPATH mode

run:

    GO111MODULE=on go get github.com/oncilla/confex@latest


## Usage

Supply any yml, json, or toml file to explore it.

    confex config.json

confex also supports piped input.

    echo $( curl http://headers.jsontest.com/ ) | confex

## Key bindings

    Basic navigation
    --------------------------------------
     k,  <UP>       Move up
     j,  <Down>     Move down
     gg, <Home>     Go to top
     G,  <End>      Go to end
     e,  <Enter>    Toggle expand
     E              Expand all
     C              Collapse all

     h              Show help message
     q, <ctrl-c>    Quit


    Extended navigation
    --------------------------------------
     <ctrl-d>       Half page down outline
     <ctrl-u>       Half page up outline
     <ctrl-f>       Page down outline
     <ctrl-b>       Page up outline


![](sample/sample.gif)

## Related Work

This project is inspired by https://github.com/gulyasm/jsonui
