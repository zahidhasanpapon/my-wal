# WAL (Write-Ahead Log)

A Write-Ahead Logging (WAL) implementation in Go.

## Project Overview

WAL is a technique used in database systems to ensure data integrity by logging changes before they
are applied to the database.

## Prerequisites

- [Bazel](https://bazel.build/install) build system
- [Go](https://golang.org/doc/install) programming language

## Building the Project

To build the project:

```bash
bazel build //:go_wal
```

## Running the Project

To run the project:

```bash
bazel run //:go_wal
```

## Development

### Update Dependencies

To update Bazel dependencies:

```bash
bazel run //:gazelle-update-repos
```

### Generate BUILD Files

To generate/update BUILD files:

```bash
bazel run //:gazelle
```
