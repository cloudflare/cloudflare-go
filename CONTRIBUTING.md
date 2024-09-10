## Setting up the environment

### Install Go 1.18+

Install go by following relevant directions [here](https://go.dev/doc/install).

## Modifying/Adding code

Most of the SDK is generated code, and any modified code will be overridden on the next generation. The
`examples/` directory is an exception and will never be overridden.

## Adding and running examples

All files in the `examples/` directory are not modified by the Stainless generator and can be freely edited or
added to.

```bash
# add an example to examples/<your-example>/main.go

package main

func main() {
  // ...
}
```

```bash
go run ./examples/<your-example>
```

## Using the repository from source

To use a local version of this library from source in another project, edit the `go.mod` with a replace
directive. This can be done through the CLI with the following:

```bash
go mod edit -replace github.com/cloudflare/cloudflare-go=/path/to/cloudflare-go
```

## Running tests

Most tests require you to [set up a mock server](https://github.com/stoplightio/prism) against the OpenAPI spec to run the tests.

```bash
# you will need npm installed
npx prism mock path/to/your/openapi.yml
```

```bash
go test ./...
```

## Formatting

This library uses the standard gofmt code formatter:

```bash
gofmt -s -w .
```
