# Logs

## Overview

Everything related to Logs

### Available Operations

* [ListLogs](#listlogs) - List the logs from a ledger

## ListLogs

List the logs from a ledger, sorted by ID in descending order.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/types"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Logs.ListLogs(ctx, operations.ListLogsRequest{
        After: formance.String("1234"),
        Cursor: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        EndTime: types.MustTimeFromString("2020-11-28T02:15:07.561Z"),
        EndTimeDeprecated: types.MustTimeFromString("2022-12-10T00:25:28.749Z"),
        Ledger: "ledger001",
        PageSize: formance.Int64(969810),
        PageSizeDeprecated: formance.Int64(666767),
        PaginationToken: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        StartTime: types.MustTimeFromString("2021-08-29T10:25:27.700Z"),
        StartTimeDeprecated: types.MustTimeFromString("2022-10-16T05:02:54.746Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LogsCursorResponse != nil {
        // handle response
    }
}
```
