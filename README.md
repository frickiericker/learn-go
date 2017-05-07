[![Build Status][travis-badge]][travis-url]

This repository contains my micro-projects for learning how to code in
[Go][golang].

[travis-badge]: https://travis-ci.org/frickiericker/learn-go.svg?branch=master
[travis-url]: https://travis-ci.org/frickiericker/learn-go
[golang]: https://golang.org

## Code organization

Assuming default environment setting (GOPATH=~/go), repository root should be
located at:

    ~/go/src/github.com/frickiericker/learn-go

Then micro-projects are laid out as follows:

    ~/go/src/github.com/frickiericker/learn-go/00-hello
    ~/go/src/github.com/frickiericker/learn-go/01-package
    ...
    ~/go/src/github.com/frickiericker/learn-go/NN-name

To run 00-hello for example:

    % cd ~/go/src/github.com/frickiericker/learn-go/00-hello
    % go test
    % go build
    % ./00-hello

## Reading materials

### Documentations

- [Official documentations](https://golang.org/doc)
  - [Language specification](https://golang.org/ref/spec)
  - [Standard library](https://golang.org/pkg)
  - [FAQ](https://golang.org/doc/faq)

### Style & code organization

- https://github.com/golang/go/wiki/CodeReviewComments
- https://golang.org/doc/code.html#Organization

### Random resources

- [Go by Example](https://gobyexample.com)

### Code to read

- https://github.com/golang/go
- https://github.com/junegunn/fzf
- https://github.com/motemen/ghq

### Books

TODO
