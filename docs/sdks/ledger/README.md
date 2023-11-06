# Ledger
(*Ledger*)

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
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
    )


    var ledger string = "ledger001"

    ctx := context.Background()
    res, err := s.Ledger.GetLedgerInfo(ctx, ledger)
    if err != nil {
        log.Fatal(err)
    }

    if res.LedgerInfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `ledger`                                              | *string*                                              | :heavy_check_mark:                                    | Name of the ledger.                                   | ledger001                                             |


### Response

**[*operations.GetLedgerInfoResponse](../../models/operations/getledgerinforesponse.md), error**

