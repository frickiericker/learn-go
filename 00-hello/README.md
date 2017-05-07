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

We can have many packages in a single repository. So let us create a directory
named NN-name for each topic. Maybe I should start with a hello world:

    mkdir learn-go/00-hello
    cd learn-go/00-hello

[codeorg]: https://golang.org/doc/code.html#Organization

## Run

    % go build
    % ./00-hello
    Hello, world!

I have created hello.go and main.go. The former defines `MakeMessage()` function
and the latter calls it. Go automatically linked these two files so that I could
call `MakeMessage()` from main.go without importing anything. Looks nice.

## Testing

hello_test.go tests hello.go. `go test` runs all tests in the current directory:

    % go test -v
    === RUN   TestMakeMessage_mustStartWithHello
    --- PASS: TestMakeMessage_mustStartWithHello (0.00s)
    === RUN   TestMakeMessage_mustEndWithPunctuation
    --- PASS: TestMakeMessage_mustEndWithPunctuation (0.00s)
    PASS
    ok      github.com/frickiericker/learn-go/00-hello      0.002s

## Documentation

    % go doc MakeMessage
    func MakeMessage() string
        MakeMessage creates a greeting message.
