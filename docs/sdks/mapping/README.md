# Mapping
(*.Mapping*)

## Overview

Everything related to Mapping

### Available Operations

* [GetMapping](#getmapping) - Get the mapping of a ledger
* [UpdateMapping](#updatemapping) - Update the mapping of a ledger

## GetMapping

Get the mapping of a ledger

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
    res, err := s.Mapping.GetMapping(ctx, ledger)
    if err != nil {
        log.Fatal(err)
    }

    if res.MappingResponse != nil {
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

**[*operations.GetMappingResponse](../../models/operations/getmappingresponse.md), error**


## UpdateMapping

Update the mapping of a ledger

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


    mapping := shared.Mapping{
        Contracts: []shared.Contract{
            shared.Contract{
                Account: formancegosdk.String("users:001"),
                Expr: shared.Expr{},
            },
        },
    }

    var ledger string = "ledger001"

    ctx := context.Background()
    res, err := s.Mapping.UpdateMapping(ctx, mapping, ledger)
    if err != nil {
        log.Fatal(err)
    }

    if res.MappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `mapping`                                             | [shared.Mapping](../../models/shared/mapping.md)      | :heavy_check_mark:                                    | N/A                                                   |                                                       |
| `ledger`                                              | *string*                                              | :heavy_check_mark:                                    | Name of the ledger.                                   | ledger001                                             |


### Response

**[*operations.UpdateMappingResponse](../../models/operations/updatemappingresponse.md), error**

