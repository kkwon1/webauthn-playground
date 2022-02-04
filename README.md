# WebAuthn Playground Service

A small project I'm working on to play around with WebAuthn concepts. This is the server side that will handle user registration and authentication.

## Description

An implementation of the concepts introduced in the [W3C specs](https://www.w3.org/TR/webauthn-2/#sctn-intro)

- Another [great resource](https://webauthn.guide/) that's friendlier on the eyes

## Prerequisites

This project can be run using docker or running the service locally on your machine. Make sure you've installed all the correct software before running the service. Please create a `.env` file in your root directory with the correct environment variables.

### Local

The service is written in Go and uses MongoDB for data persistence. You will need to install both onto your machine.

- You can find the official binary releases [here](https://golang.org/dl/)
- Find the correct installation steps for your OS [here](https://golang.org/doc/install)

This project is using go modules. After pulling the repo, simply run

```bash
go get -v -d ./...
```

## Usage

### Local

From the root directory you can build and run the binary file

```bash
go build main.go
main.exe
```

Or if you'd like to directly run the service without building

```bash
go run main.go
```

## Running Tests

To run all tests

```bash
go test -v ./...
```

To run a single test file, specify the test file path

```bash
go test -v ./test/dir/<FILE_NAME>
```

## Linting

You can also optionally have golangci-lint installed on your computer. You can check out their [github page](https://github.com/golangci/golangci-lint)

```bash
golangci-lint run
```

## Authors

- **Kevin Kwon** - [portfolio](https://kkwon1.github.io/portfolio/)
