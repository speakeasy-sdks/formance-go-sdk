# Transactions
(*Transactions*)

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
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Transactions.CreateTransactions(ctx, operations.CreateTransactionsRequest{
        Transactions: shared.Transactions{
            Transactions: []shared.TransactionData{
                shared.TransactionData{
                    Metadata: map[string]interface{}{
                        "admin": "string",
                        "a": "string",
                    },
                    Postings: []shared.Posting{
                        shared.Posting{
                            Amount: 100,
                            Asset: "COIN",
                            Destination: "users:002",
                            Source: "users:001",
                        },
                    },
                    Reference: formancegosdk.String("ref:001"),
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
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Transactions.AddMetadataOnTransaction(ctx, operations.AddMetadataOnTransactionRequest{
        RequestBody: map[string]interface{}{
            "admin": "string",
            "a": "string",
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
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Transactions.CountTransactions(ctx, operations.CountTransactionsRequest{
        Account: formancegosdk.String("users:001"),
        Destination: formancegosdk.String("users:001"),
        Ledger: "ledger001",
        Metadata: &operations.CountTransactionsMetadata{},
        Reference: formancegosdk.String("ref:001"),
        Source: formancegosdk.String("users:001"),
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
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Transactions.CreateTransaction(ctx, operations.CreateTransactionRequest{
        PostTransaction: shared.PostTransaction{
            Metadata: map[string]interface{}{
                "admin": "string",
                "a": "string",
            },
            Postings: []shared.Posting{
                shared.Posting{
                    Amount: 100,
                    Asset: "COIN",
                    Destination: "users:002",
                    Source: "users:001",
                },
            },
            Reference: formancegosdk.String("ref:001"),
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
        },
        Ledger: "ledger001",
        Preview: formancegosdk.Bool(true),
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
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
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
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Transactions.ListTransactions(ctx, operations.ListTransactionsRequest{
        Account: formancegosdk.String("users:001"),
        After: formancegosdk.String("1234"),
        Cursor: formancegosdk.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Destination: formancegosdk.String("users:001"),
        Ledger: "ledger001",
        Metadata: &operations.ListTransactionsMetadata{},
        PaginationToken: formancegosdk.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Reference: formancegosdk.String("ref:001"),
        Source: formancegosdk.String("users:001"),
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
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
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

