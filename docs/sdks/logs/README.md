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
        EndTime: types.MustTimeFromString("2022-03-15T07:22:15.330Z"),
        EndTimeDeprecated: types.MustTimeFromString("2021-08-13T16:19:19.906Z"),
        Ledger: "ledger001",
        PageSize: formancegosdk.Int64(520478),
        PageSizeDeprecated: formancegosdk.Int64(780529),
        PaginationToken: formancegosdk.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        StartTime: types.MustTimeFromString("2022-10-06T15:49:54.663Z"),
        StartTimeDeprecated: types.MustTimeFromString("2021-09-20T20:35:01.256Z"),
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

