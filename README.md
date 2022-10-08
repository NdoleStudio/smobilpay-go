# Smobilpay Go SDK

[![Build](https://github.com/NdoleStudio/smobilpay-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/smobilpay-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/smobilpay-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/smobilpay-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/smobilpay-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/smobilpay-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/smobilpay-go)](https://goreportcard.com/report/github.com/NdoleStudio/smobilpay-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/smobilpay-go)](https://github.com/NdoleStudio/smobilpay-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/smobilpay-go?color=brightgreen)](https://github.com/NdoleStudio/smobilpay-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/smobilpay-go)](https://pkg.go.dev/github.com/NdoleStudio/smobilpay-go)


This package provides an unofficial `go` client for the [Smobilpay API](https://apidocs.smobilpay.com/s3papi/index.html)

## Installation

`smobilpay-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/smobilpay-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/smobilpay-go"
```


## Implemented

- [Status Codes](#status-codes)
    - `GET /200`: OK

## Usage

### Initializing the Client

An instance of the client can be created using `New()`.

```go
package main

import (
	"github.com/NdoleStudio/smobilpay-go"
)

func main()  {
	client := smobilpay.New(smobilpay.WithDelay(200))
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
status, response, err := statusClient.Status.Ok(context.Background())
if err != nil {
    //handle error
}
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
