# Search
(*Search*)

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
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Search.Search(ctx, shared.Query{
        After: []string{
            "u",
            "s",
            "e",
            "r",
            "s",
            ":",
            "0",
            "0",
            "2",
        },
        Cursor: formancegosdk.String("YXVsdCBhbmQgYSBtYXhpbXVtIG1heF9yZXN1bHRzLol="),
        Ledgers: []string{
            "q",
            "u",
            "i",
            "c",
            "k",
            "s",
            "t",
            "a",
            "r",
            "t",
        },
        Policy: formancegosdk.String("OR"),
        Sort: formancegosdk.String("txid:asc"),
        Terms: []string{
            "d",
            "e",
            "s",
            "t",
            "i",
            "n",
            "a",
            "t",
            "i",
            "o",
            "n",
            "=",
            "c",
            "e",
            "n",
            "t",
            "r",
            "a",
            "l",
            "_",
            "b",
            "a",
            "n",
            "k",
            "1",
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
| `request`                                             | [shared.Query](../../pkg/models/shared/query.md)      | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.SearchResponse](../../pkg/models/operations/searchresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
