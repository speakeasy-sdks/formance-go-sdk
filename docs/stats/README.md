# Stats

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
	"github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ReadStatsRequest{
        Ledger: "ledger001",
    }

    res, err := s.Stats.ReadStats(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatsResponse != nil {
        // handle response
    }
}
```
