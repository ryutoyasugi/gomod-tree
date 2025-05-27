# gomod-tree
[![Test](https://github.com/ryutoyasugi/gomod-tree/actions/workflows/test.yml/badge.svg)](https://github.com/ryutoyasugi/gomod-tree/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ryutoyasugi/gomod-tree)](https://goreportcard.com/report/github.com/ryutoyasugi/gomod-tree)
[![Version](https://img.shields.io/github/v/tag/ryutoyasugi/gomod-tree)](https://github.com/ryutoyasugi/gomod-tree/releases/latest)

Print `go mod` dependency tree.  
This command call `go mod graph` internally.

## Install
```sh
$ go install github.com/ryutoyasugi/gomod-tree@latest
```

## Usage
```sh
$ gomod-tree
- github.com/ryutoyasugi/gomod-tree
    - github.com/inconshreveable/mousetrap@v1.0.0
    - github.com/spf13/cobra@v1.4.0
        - github.com/cpuguy83/go-md2man/v2@v2.0.1
            - github.com/russross/blackfriday/v2@v2.1.0
        - github.com/inconshreveable/mousetrap@v1.0.0
        - github.com/spf13/pflag@v1.0.5
        - gopkg.in/yaml.v2@v2.4.0
            - gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
    - github.com/spf13/pflag@v1.0.5
```
Specify depth (default 3)
```sh
$ gomod-tree -d 1
- github.com/ryutoyasugi/gomod-tree
    - github.com/inconshreveable/mousetrap@v1.0.0
    - github.com/spf13/cobra@v1.4.0
    - github.com/spf13/pflag@v1.0.5
```
