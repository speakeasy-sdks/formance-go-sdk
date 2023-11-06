# Stats
(*Stats*)

## Overview

Everything related to Stats

### Available Operations

* [ReadStats](#readstats) - Get statistics from a ledger

## ReadStats

Get statistics from a ledger. (aggregate metrics on accounts and transactions)


### Example Usage

```go
package main

import(
	"context"
	"log"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
    )


    var ledger string = "ledger001"

    ctx := context.Background()
    res, err := s.Stats.ReadStats(ctx, ledger)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `ledger`                                              | *string*                                              | :heavy_check_mark:                                    | name of the ledger                                    | ledger001                                             |


### Response

**[*operations.ReadStatsResponse](../../models/operations/readstatsresponse.md), error**

