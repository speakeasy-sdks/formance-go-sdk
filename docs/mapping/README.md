# Mapping

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
    res, err := s.Mapping.GetMapping(ctx, operations.GetMappingRequest{
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MappingResponse != nil {
        // handle response
    }
}
```

## UpdateMapping

Update the mapping of a ledger

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Mapping.UpdateMapping(ctx, operations.UpdateMappingRequest{
        Mapping: shared.Mapping{
            Contracts: []shared.Contract{
                shared.Contract{
                    Account: formance.String("users:001"),
                    Expr: map[string]interface{}{
                        "nobis": "enim",
                    },
                },
                shared.Contract{
                    Account: formance.String("users:001"),
                    Expr: map[string]interface{}{
                        "nemo": "minima",
                        "excepturi": "accusantium",
                        "iure": "culpa",
                    },
                },
            },
        },
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MappingResponse != nil {
        // handle response
    }
}
```
