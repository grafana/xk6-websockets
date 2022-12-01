# xk6-websockets

This extension adds a PoC [Websockets API](https://websockets.spec.whatwg.org) implementation to [k6](https://www.k6.io).

This is meant to try to implement the specification as close as possible without doing stuff that don't make sense in k6 like:

1. not reporting errors
2. not allowing some ports and other security workarounds
3. supporting Blob as message

It supports additional k6 specific features such as:

* Custom metrics tags
* Cookie jar
* Headers customization
* Support for ping/pong which isn't part of the specification
* Compression Support (The only supported algorithm currently is `deflate`)

It is implemented using the [xk6](https://k6.io/blog/extending-k6-with-xk6/) system.

## Requirements

* [Golang 1.17+](https://go.dev/)
* [Git](https://git-scm.com/)
* [xk6](https://github.com/grafana/xk6) (`go install go.k6.io/xk6/cmd/xk6@latest`)

## Getting started  

1. Build the k6's binary:

  ```shell
  $ make build
  ```

2. Run an example:

  ```shell
  $ ./k6 run ./examples/test-api.k6.io.js
  ```

## Discrepancies with the specifications

* binaryType is "ArrayBuffer" by default instead of "Blob" and will throw an exception if it's tried to be changed as "Blob" is not supported by k6.
