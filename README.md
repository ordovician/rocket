# Go Rocket

A toy project to simulate rocket launches in the Go programming language. It is intended to demonstrate the usage of modules, packages, using of tests as well as how embedding works in Go.

## Running tests
Each of the sub packages have a number of tests which can be run. You can run the tests in the github.com/ordovician/rocket package like this. Notice the `-v` switch. It shows you what tests actually got run.

    ❯ go test -v .


We can also test individual sub-packages like this:

    ❯ go test -v ./physics
    ❯ go test -v ./engine

## Running executables
There are actual commands you can run in the cmd sub directory such as launcher, playground and server. Not all of these are finnished development.

You can run the cmd/server command which demonstrates a simple Web service which allows you to check for engines at http://localhost:8080/engines/ or a specific engine by looking  at http://localhost:8080/engines/BE-4. 

    ❯ go run ./cmd/server
    
The point of this is to demonstrate the use of JSON parsing and generation and how to respond to HTTP requests.
