A package is a directory containing go source files. Here I modified 00-hello
so that hello.go is in a new package named `hello`.

## Imports

Imports are relative to GOPATH (~/go). So this

    import (
        "github.com/frickiericker/learn-go/01-package/hello"
    )

imports a package **locally** located at
~/go/src/github.com/frickiericker/learn-go/01-package/hello.

I could import the same package by relative path "./hello" from main.go, but
[full path is considered the best practice][fullpath].

[fullpath]: http://stackoverflow.com/a/10688069/5266681

## Testing

    % go test -v ./...
    ?       github.com/frickiericker/learn-go/01-package    [no test files]
    === RUN   TestMakeMessage_mustStartWithHello
    --- PASS: TestMakeMessage_mustStartWithHello (0.00s)
    === RUN   TestMakeMessage_mustEndWithPunctuation
    --- PASS: TestMakeMessage_mustEndWithPunctuation (0.00s)
    PASS
    ok      github.com/frickiericker/learn-go/01-package/hello      0.002s

The three dots [recursively expands to all subdirectories][dots]. For this
project it is equivalent to

    % go test -v ./hello

[dots]: http://stackoverflow.com/q/28031603/5266681

## Documentation

    % go doc hello.MakeMessage
    package hello // import "github.com/frickiericker/learn-go/01-package/hello"
    
    func MakeMessage() string
        MakeMessage creates a greeting message.
