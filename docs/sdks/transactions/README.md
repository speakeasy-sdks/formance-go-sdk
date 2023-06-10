# Transactions

## Overview

Everything related to Transactions

### Available Operations

* [CreateTransactions](#createtransactions) - Create a new batch of transactions to a ledger
* [AddMetadataOnTransaction](#addmetadataontransaction) - Set the metadata of a transaction by its ID
* [CountTransactions](#counttransactions) - Count the transactions from a ledger
* [CreateTransaction](#createtransaction) - Create a new transaction to a ledger
* [GetTransaction](#gettransaction) - Get transaction from a ledger by its ID
* [ListTransactions](#listtransactions) - List transactions from a ledger
* [RevertTransaction](#reverttransaction) - Revert a ledger transaction by its ID

## CreateTransactions

Create a new batch of transactions to a ledger

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/types"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.CreateTransactions(ctx, operations.CreateTransactionsRequest{
        Transactions: shared.Transactions{
            Transactions: []shared.TransactionData{
                shared.TransactionData{
                    Metadata: map[string]interface{}{
                        "commodi": "repudiandae",
                        "quae": "ipsum",
                        "quidem": "molestias",
                        "excepturi": "pariatur",
                    },
                    Postings: []shared.Posting{
                        shared.Posting{
                            Amount: 100,
                            Asset: "COIN",
                            Destination: "users:002",
                            Source: "users:001",
                        },
                        shared.Posting{
                            Amount: 100,
                            Asset: "COIN",
                            Destination: "users:002",
                            Source: "users:001",
                        },
                    },
                    Reference: formance.String("ref:001"),
                    Timestamp: types.MustTimeFromString("2021-12-15T00:41:38.329Z"),
                },
                shared.TransactionData{
                    Metadata: map[string]interface{}{
                        "quasi": "repudiandae",
                        "sint": "veritatis",
                        "itaque": "incidunt",
                        "enim": "consequatur",
                    },
                    Postings: []shared.Posting{
                        shared.Posting{
                            Amount: 100,
                            Asset: "COIN",
                            Destination: "users:002",
                            Source: "users:001",
                        },
                        shared.Posting{
                            Amount: 100,
                            Asset: "COIN",
                            Destination: "users:002",
                            Source: "users:001",
                        },
                        shared.Posting{
                            Amount: 100,
                            Asset: "COIN",
                            Destination: "users:002",
                            Source: "users:001",
                        },
                    },
                    Reference: formance.String("ref:001"),
                    Timestamp: types.MustTimeFromString("2022-08-09T16:21:07.003Z"),
                },
            },
        },
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateTransactionsRequest](../../models/operations/createtransactionsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CreateTransactionsResponse](../../models/operations/createtransactionsresponse.md), error**


## AddMetadataOnTransaction

Set the metadata of a transaction by its ID

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
    res, err := s.Transactions.AddMetadataOnTransaction(ctx, operations.AddMetadataOnTransactionRequest{
        RequestBody: map[string]interface{}{
            "distinctio": "quibusdam",
            "labore": "modi",
            "qui": "aliquid",
        },
        Ledger: "ledger001",
        Txid: 1234,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AddMetadataOnTransactionRequest](../../models/operations/addmetadataontransactionrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.AddMetadataOnTransactionResponse](../../models/operations/addmetadataontransactionresponse.md), error**


## CountTransactions

Count the transactions from a ledger

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
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.CountTransactions(ctx, operations.CountTransactionsRequest{
        Account: formance.String("users:001"),
        Destination: formance.String("users:001"),
        EndTime: types.MustTimeFromString("2021-11-23T10:34:02.904Z"),
        EndTimeDeprecated: types.MustTimeFromString("2022-11-01T19:07:16.800Z"),
        Ledger: "ledger001",
        Metadata: &operations.CountTransactionsMetadata{},
        Reference: formance.String("ref:001"),
        Source: formance.String("users:001"),
        StartTime: types.MustTimeFromString("2021-11-22T01:26:35.048Z"),
        StartTimeDeprecated: types.MustTimeFromString("2022-11-08T13:10:11.700Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CountTransactionsRequest](../../models/operations/counttransactionsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CountTransactionsResponse](../../models/operations/counttransactionsresponse.md), error**


## CreateTransaction

Create a new transaction to a ledger

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/types"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.CreateTransaction(ctx, operations.CreateTransactionRequest{
        PostTransaction: shared.PostTransaction{
            Metadata: map[string]interface{}{
                "excepturi": "tempora",
                "facilis": "tempore",
                "labore": "delectus",
            },
            Postings: []shared.Posting{
                shared.Posting{
                    Amount: 100,
                    Asset: "COIN",
                    Destination: "users:002",
                    Source: "users:001",
                },
                shared.Posting{
                    Amount: 100,
                    Asset: "COIN",
                    Destination: "users:002",
                    Source: "users:001",
                },
            },
            Reference: formance.String("ref:001"),
            Script: &shared.PostTransactionScript{
                Plain: "vars {
            account $user
            }
            send [COIN 10] (
            	source = @world
            	destination = $user
            )
            ",
                Vars: &shared.PostTransactionScriptVars{},
            },
            Timestamp: types.MustTimeFromString("2022-03-31T00:30:19.135Z"),
        },
        Ledger: "ledger001",
        Preview: formance.Bool(true),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateTransactionRequest](../../models/operations/createtransactionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateTransactionResponse](../../models/operations/createtransactionresponse.md), error**


## GetTransaction

Get transaction from a ledger by its ID

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
    res, err := s.Transactions.GetTransaction(ctx, operations.GetTransactionRequest{
        Ledger: "ledger001",
        Txid: 1234,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetTransactionRequest](../../models/operations/gettransactionrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetTransactionResponse](../../models/operations/gettransactionresponse.md), error**


## ListTransactions

List transactions from a ledger, sorted by txid in descending order.

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
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.ListTransactions(ctx, operations.ListTransactionsRequest{
        Account: formance.String("users:001"),
        After: formance.String("1234"),
        Cursor: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Destination: formance.String("users:001"),
        EndTime: types.MustTimeFromString("2022-03-17T20:21:28.792Z"),
        EndTimeDeprecated: types.MustTimeFromString("2021-03-17T21:24:26.606Z"),
        Ledger: "ledger001",
        Metadata: &operations.ListTransactionsMetadata{},
        PageSize: formance.Int64(572252),
        PageSizeDeprecated: formance.Int64(638921),
        PaginationToken: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Reference: formance.String("ref:001"),
        Source: formance.String("users:001"),
        StartTime: types.MustTimeFromString("2022-02-09T13:58:59.361Z"),
        StartTimeDeprecated: types.MustTimeFromString("2020-12-17T08:08:22.040Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListTransactionsRequest](../../models/operations/listtransactionsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListTransactionsResponse](../../models/operations/listtransactionsresponse.md), error**


## RevertTransaction

Revert a ledger transaction by its ID

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
    res, err := s.Transactions.RevertTransaction(ctx, operations.RevertTransactionRequest{
        Ledger: "ledger001",
        Txid: 1234,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RevertTransactionRequest](../../models/operations/reverttransactionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.RevertTransactionResponse](../../models/operations/reverttransactionresponse.md), error**

