# Development Guide

This document contains notes about development and testing of Protocol Buffers Language Server.

## Prerequisites

- [Go 1.13](https://golang.org/dl) (for development)
- [Bazelisk](https://github.com/bazelbuild/bazelisk) (explained at [Build and Test](#build-and-test))

## Build and Test

This project uses [Bazel](https://bazel.build) and also Bazelisk for build and test.
Bazelisk installs Bazel versioned by `.bazelversion` if not installed yet, and uses Bazel of the version.
So all you need is to install Bazelisk.

```
$ go get github.com/bazelbuild/bazelisk
```

And Bazel controls the versions of Go, Protocol Buffers and something like that.
Thus, you don't need to care about their versions and install them to build or test.
To develop you need to use Go 1.13.

To build this project, run the following command.
This builds it with Bazel.

```
$ make build
```

To test this project, run the following command.
This tests it with Bazel.

```
$ make test
```

To update Go Modules, run the following command.
This updates `go.mod` and `go.sum` and then updates the related Bazel files.

```
$ make dep
```

The generated files, like `.mock.go`, are not controlled by Git.
And Bazel generates the files but the stuff is put into Bazel's sandbox.
As a result, the files don't appear in your editor or IDE.

To link generated files to your workspace and then make them appear in your editor or IDE, run the following command.

```
$ make expose-generated-go
```

## Debug

This project provides some flags for debugging.
To know the way to debug, run it with `--help`.
