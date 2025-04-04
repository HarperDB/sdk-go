# Harper SDK for Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/HarperDB/sdk-go)](https://pkg.go.dev/github.com/HarperDB/sdk-go)

This is the Go SDK for [Harper](https://harpersystems.dev/).

## Requirements

- >= Go 1.18

## Installation

```
go get github.com/HarperDB/sdk-go
```

## Quickstart

```go
client := harper.NewClient("http://localhost:9925", "HDB_ADMIN", "password")
client.CreateSchema("dog")
```
