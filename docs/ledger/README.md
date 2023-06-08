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
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.GetLedgerInfo(ctx, operations.GetLedgerInfoRequest{
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LedgerInfoResponse != nil {
        // handle response
    }
}
```
