# Logs
(*Logs*)

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
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/types"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(shared.Security{
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Logs.ListLogs(ctx, operations.ListLogsRequest{
        After: formancegosdk.String("1234"),
        Cursor: formancegosdk.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        EndTime: types.MustTimeFromString("2022-09-04T05:10:40.788Z"),
        EndTimeDeprecated: types.MustTimeFromString("2022-12-21T16:19:06.983Z"),
        Ledger: "ledger001",
        PageSize: formancegosdk.Int64(279551),
        PageSizeDeprecated: formancegosdk.Int64(61287),
        PaginationToken: formancegosdk.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        StartTime: types.MustTimeFromString("2021-01-03T01:50:13.003Z"),
        StartTimeDeprecated: types.MustTimeFromString("2022-01-22T19:40:11.309Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LogsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.ListLogsRequest](../../models/operations/listlogsrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.ListLogsResponse](../../models/operations/listlogsresponse.md), error**

