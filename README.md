# hoover [![Build Status](https://travis-ci.org/phedoreanu/hoover.svg?branch=master)](https://travis-ci.org/phedoreanu/hoover) [![Coverage Status](https://coveralls.io/repos/github/phedoreanu/hoover/badge.svg)](https://coveralls.io/github/phedoreanu/hoover) [![Code Climate](https://codeclimate.com/github/phedoreanu/hoover/badges/gpa.svg)](https://codeclimate.com/github/phedoreanu/hoover) [![Go Report Card](https://goreportcard.com/badge/github.com/phedoreanu/hoover)](https://goreportcard.com/report/github.com/phedoreanu/hoover)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org) [![GoDoc](https://godoc.org/github.com/phedoreanu/hoover?status.svg)](https://godoc.org/github.com/phedoreanu/hoover)

Hoover is a little Go room cleaning library.

### Installation
Hoover requires Go 1.5 or later.

This repo contains just the _library_, check out [vacuum](https://github.com/phedoreanu/vacuum)/[godocs](https://godoc.org/github.com/phedoreanu/hoover#ex-package) for an example.

```
go get -u github.com/phedoreanu/vacuum
```

### Usage
Invoke `vacuum` in the same directory as `input.txt`.

### Testing
```
$ go test github.com/phedoreanu/hoover
ok      github.com/phedoreanu/hoover    0.005s
```

### Cyclomatic complexity
```
$ mccabe-cyclomatic -p github.com/phedoreanu/hoover
6
```

```
$ gocyclo -top 5 -avg .
3 hoover (*Patch).Update hoover.go:31:1
3 hoover NewHoover hoover.go:83:1
2 hoover (*Hoover).Vacuum hoover.go:58:1
2 hoover (*Hoover).placeDirtyPatches hoover.go:73:1
1 hoover parseUInt16 hoover.go:52:1
Average: 1.46
```
