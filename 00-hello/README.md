How to write a go program: Code organization, testing and documentation.

## Set up

[Documentation][codeorg] says that you should stick with a single workspace
pointed to by GOPATH (which defaults to ~/go) and put all projects there.

I'd just go with the default:

    mkdir ~/go
    echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshenv

And create a directory for this repository:

    mkdir -p ~/go/src/github.com/frickiericker
    cd ~/go/src/github.com/frickiericker
    git clone git@github.com:frickiericker/learn-go.git

And make a directory for this hello world project:

    mkdir ~/go/src/github.com/frickiericker/learn-go/00-hello
    cd ~/go/src/github.com/frickiericker/learn-go/00-hello

[codeorg]: https://golang.org/doc/code.html#Organization

## Imports

Imports are relative to GOPATH. So this

    import (
        "github.com/frickiericker/learn-go/00-hello/hello"
    )

imports a package **locally** located at
~/go/src/github.com/frickiericker/learn-go/00-hello/hello.

I could import the same package by relative path "./hello" from main.go
but [full path is considered the best practice][fullpath].

[fullpath]: http://stackoverflow.com/a/10688069/5266681

## Testing

    go test ./...

The three dots [recursively expands to all subdirectories][dots]. For this
hello world project it is equivalent to

    go test ./hello

[dots]: http://stackoverflow.com/q/28031603/5266681

## Running

    go run main.go

or

    go build
    ./00-hello

## Documentation

    go doc hello.MakeMessage
