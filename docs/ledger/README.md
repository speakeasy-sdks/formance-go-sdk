# Ledger

## Overview

Everything related to Ledger

### Available Operations

* [GetLedgerInfo](#getledgerinfo) - Get information about a ledger

## GetLedgerInfo

Get information about a ledger

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
    req := operations.GetLedgerInfoRequest{
        Ledger: "ledger001",
    }

    res, err := s.Ledger.GetLedgerInfo(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.LedgerInfoResponse != nil {
        // handle response
    }
}
```
