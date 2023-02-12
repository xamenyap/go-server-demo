# go-server-demo

A demonstration of how to establish an HTTP server in Go, and how to gracefully shut it down.

## Usage
To build the graceful server, simply run
```sh
$ go build -o graceful-go-server graceful/main.go
```
Then run the server
```sh
$ ./graceful-go-server
```

Alternatively you can build and run the default go server to compare the behavior.
