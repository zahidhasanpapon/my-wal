# WAL (Write-Ahead Log)

[![Go Report Card](https://goreportcard.com/badge/github.com/zahidhasanpapon/my-wal)](https://goreportcard.com/report/github.com/zahidhasanpapon/my-wal)

A high-performance, thread-safe Write-Ahead Logging (WAL) implementation in Go, designed for
building reliable storage systems.

## Prerequisites

- [Go](https://golang.org/doc/install) 1.24.3 or later
- [Bazel](https://bazel.build/install) for building (optional, Go modules also supported)
- [Protocol Buffer](https://grpc.io/docs/protoc-installation/) compiler (for regenerating protos)

## Building and Testing

### Building with Bazel

```bash
# Build everything
bazel build //...

# Run tests
bazel test //...

# Build and run a specific target
bazel run //cmd/wal:main
```

### Running Tests

```bash
# Run all tests
bazel test //...

# Run specific test
bazel test //internal/wal/writer:go_default_test
```

### Generate BUILD Files and Update Dependencies

To (re-)generate BUILD files:

```bash
bazel run //:gazelle
```
