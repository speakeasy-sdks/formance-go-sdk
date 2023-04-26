# Search

## Overview

Everything related to Search

### Available Operations

* [Search](#search) - Search

## Search

ElasticSearch query engine

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.Query{
        After: []string{
            "users:002",
            "users:002",
            "users:002",
        },
        Cursor: formance.String("YXVsdCBhbmQgYSBtYXhpbXVtIG1heF9yZXN1bHRzLol="),
        Ledgers: []string{
            "quickstart",
            "quickstart",
            "quickstart",
        },
        PageSize: formance.Int64(20107),
        Policy: formance.String("OR"),
        Sort: formance.String("txid:asc"),
        Target: formance.String("magni"),
        Terms: []string{
            "destination=central_bank1",
            "destination=central_bank1",
            "destination=central_bank1",
            "destination=central_bank1",
        },
    }

    res, err := s.Search.Search(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Response != nil {
        // handle response
    }
}
```
