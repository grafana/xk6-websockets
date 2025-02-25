# xk6-websockets

> [!WARNING]
> The `xk6-websockets` extension [has been merged](https://github.com/grafana/k6/pull/4131) to the [main k6 repository](https://github.com/grafana/k6). Please contribute and [open issues there](https://github.com/grafana/k6/issues). This repository is no longer maintained.

This extension adds a PoC [Websockets API](https://websockets.spec.whatwg.org) implementation to [k6](https://www.k6.io).

This is meant to try to implement the specification as close as possible without doing stuff that don't make sense in k6 like:

1. not reporting errors
2. not allowing some ports and other security workarounds

It supports additional k6 specific features such as:

* Custom metrics tags
* Cookie jar
* Headers customization
* Support for ping/pong which isn't part of the specification
* Compression Support (The only supported algorithm currently is `deflate`)

It is implemented using the [xk6](https://k6.io/blog/extending-k6-with-xk6/) system.

## Requirements

* [Golang 1.19+](https://go.dev/)
* [Git](https://git-scm.com/)
* [xk6](https://github.com/grafana/xk6) (`go install go.k6.io/xk6/cmd/xk6@latest`)
* [curl](https://curl.se/) (downloading the k6 core's linter rule-set)

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

* `binaryType` does not have a default value (in contrast to the spec, [which suggests `"blob"` as default](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/binaryType)),
so in order to successfully receive binary messages a `binaryType` must be explicitly set either to `"arraybuffer"` (for `ArrayBuffer`)
or `"blob"` (for `Blob`).

## Contributing

Contributing to this repository is following general k6's [contribution guidelines](https://github.com/grafana/k6/blob/master/CONTRIBUTING.md) since the long-term goal is to merge this extension into the main k6 repository.

### Testing

To run the test you can use the `make test` target.

### Linting

To run the linter you can use the `make lint` target.

> [!IMPORTANT]  
> By default there is golangci-lint config presented. Since the long-term goal is to merge the module back to the grafana/k6 we use the k6's linter rules. The rule set will be downloaded automatically while the first run of the `make lint` or you could do that manually by running `make linter-config`.
