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
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.Search(ctx, shared.Query{
        After: []string{
            "users:002",
        },
        Cursor: formance.String("YXVsdCBhbmQgYSBtYXhpbXVtIG1heF9yZXN1bHRzLol="),
        Ledgers: []string{
            "quickstart",
        },
        PageSize: formance.Int64(666767),
        Policy: formance.String("OR"),
        Sort: formance.String("txid:asc"),
        Target: formance.String("mollitia"),
        Terms: []string{
            "destination=central_bank1",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Response != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Query](../../models/shared/query.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.SearchResponse](../../models/operations/searchresponse.md), error**

