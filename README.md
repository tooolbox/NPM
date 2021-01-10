# Fiserv Go SDK

[![Godoc](https://pkg.go.dev/github.com/tooolbox/fiserv?status.svg)](https://godoc.org/github.com/tooolbox/fiserv)

Fiserv RESTful Web Services

Payment Gateway API Specification.
- API version: 6.14.0.20201015.001

## Usage

With Go Modules: 

```go
package main

import (
    "github.com/tooolbox/fiserv/simple"
)

func main() {
    cfg := simple.Config{
        ApiKey: "...",
        ApiSecret: "...",
        Environment: simple.ProdEnvironment,
    }

    gw, err := simple.NewGateway(cfg)
    if err != nil {
        panic(err)
    }

    // ...
}
```

## Developing

To re-generate the OpenAPI types and client:

```sh
> cd tools
> go generate # install the tools
> cd ..
> go generate # generate the code
```

## Reference Documentation

https://docs.firstdata.com/org/gateway/docs/api
