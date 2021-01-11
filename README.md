# Fiserv Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/tooolbox/fiserv.svg)](https://pkg.go.dev/github.com/tooolbox/fiserv)

Fiserv RESTful Web Services

Payment Gateway API Specification.
- API version: 6.14.0.20201015.001

## Usage

With Go Modules: 

```go
package main

import (
    "fmt"
    "github.com/tooolbox/fiserv"
)

func main() {
    creds := fiserv.Credentials{
        ApiKey:    "...",
        ApiSecret: "...",
    }

    clnt, err := fiserv.NewClientWithResponses(fiserv.SandboxServer, creds)
    if err != nil {
        panic(err)
    }

    // Calls will automatically generate the authentication signature.
    resp, err := clnt.SubmitPrimaryTransactionWithResponse(
        context.Background(),
        &fiserv.SubmitPrimaryTransactionParams{
            // Params will be automatically populated
            // with signature-related headers.
        },
        &fiserv.SubmitPrimaryTransactionJSONBody{
            // ...
        },
    )
    if err != nil {
        panic(err)
    }

    fmt.Println(resp.JSON200.ApprovedAmount.Total)
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
