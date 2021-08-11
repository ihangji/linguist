# linguist

[![Go Report Card](https://goreportcard.com/badge/github.com/ihangji/linguist)](https://goreportcard.com/report/github.com/ihangji/linguist)
[![GoDoc](https://godoc.org/github.com/ihangji/linguist?status.svg)](https://godoc.org/github.com/ihangji/linguist)

Go port of github linguist.

## Overview

Linguist is a package that detects the language through the files or directories.

## Install

1. To install linguist package, you need to install Go and set your Go workspace first.

```shell
go get -u github.com/ihangji/linguist
```

2. Import linguist package.

```Go
import github.com/ihangji/linguist
```

## Quick start

Detecting the language through the files.

```Go
package main

import github.com/ihangji/linguist

func main() {
    appPythonfile := filepath.Join("Testdata", "app-python", "app.py")
    l, err := ProcessFile(appPythonfile)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Printf("language: %s", l)
    // Output: language: Python
}
```

Detecting  the language list of the directory.

```Go
package main

import github.com/ihangji/linguist

func main() {
    var sortResult sortableResult
    appPath := filepath.Join("Testdata", "app-python")
    lanuages, err := ProcessDir(appPath)
    if err != nil {
        log.Fatalln(err)
    }
    sortResult = lanuages
    sort.Sort(sortResult)
    for _, i := range sortResult {
        fmt.Printf("language: %s, Percent: %f, Color: %s\n",
            i.Language, i.Percent, i.Color)
    }
}
```

## Contibuting code

Welcome to contribute your code.

To update to the latest version of linguist, run

```shell
git clone https://github.com/github/linguist data/linguist
go generate .
go generate ./data
rm -rf data/linguist
```
