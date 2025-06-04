# Go chanutils package

[![Go Reference](https://pkg.go.dev/badge/github.com/sinu5oid/chanutils.svg)](https://pkg.go.dev/github.com/sinu5oid/chanutils)

This repository contains a distributable package with typed (using generics) cache wrappers

## Install the package

```
$ go get gituhb.com/sinu5oid/chanutils
```

## Contents

* [Merge](merge.go) - a function that merges multiple homogenous channels into one
* [AnyError](anyerror.go) - a function that returns first non-nil error from provided channels (if
  any), or nil otherwise

## Clone the project

```
$ git clone https://github.com/sinu5oid/chanutils
$ cd chanutils
```
