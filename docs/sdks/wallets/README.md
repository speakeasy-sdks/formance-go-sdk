# Wallets
(*Wallets*)

## Overview

Everything related to Wallets

### Available Operations

* [ConfirmHold](#confirmhold) - Confirm a hold
* [CreateBalance](#createbalance) - Create a balance
* [CreateWallet](#createwallet) - Create a new wallet
* [CreditWallet](#creditwallet) - Credit a wallet
* [DebitWallet](#debitwallet) - Debit a wallet
* [GetBalance](#getbalance) - Get detailed balance
* [GetHold](#gethold) - Get a hold
* [GetHolds](#getholds) - Get all holds for a wallet
* [GetTransactions](#gettransactions)
* [GetWallet](#getwallet) - Get a wallet
* [ListBalances](#listbalances) - List balances of a wallet
* [ListWallets](#listwallets) - List all wallets
* [UpdateWallet](#updatewallet) - Update a wallet
* [VoidHold](#voidhold) - Cancel a hold
* [WalletsgetServerInfo](#walletsgetserverinfo) - Get server info

## ConfirmHold

Confirm a hold

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


    var holdID string = "string"

    confirmHoldRequest := &shared.ConfirmHoldRequest{
        Amount: formancegosdk.Int64(100),
        Final: formancegosdk.Bool(true),
    }

    ctx := context.Background()
    res, err := s.Wallets.ConfirmHold(ctx, holdID, confirmHoldRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |
| `holdID`                                                                | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `confirmHoldRequest`                                                    | [*shared.ConfirmHoldRequest](../../models/shared/confirmholdrequest.md) | :heavy_minus_sign:                                                      | N/A                                                                     |


### Response

**[*operations.ConfirmHoldResponse](../../models/operations/confirmholdresponse.md), error**


## CreateBalance

Create a balance

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


    var id string = "string"

    createBalanceRequest := &shared.CreateBalanceRequest{
        Name: "string",
    }

    ctx := context.Background()
    res, err := s.Wallets.CreateBalance(ctx, id, createBalanceRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBalanceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ctx`                                                                       | [context.Context](https://pkg.go.dev/context#Context)                       | :heavy_check_mark:                                                          | The context to use for the request.                                         |
| `id`                                                                        | *string*                                                                    | :heavy_check_mark:                                                          | N/A                                                                         |
| `createBalanceRequest`                                                      | [*shared.CreateBalanceRequest](../../models/shared/createbalancerequest.md) | :heavy_minus_sign:                                                          | N/A                                                                         |


### Response

**[*operations.CreateBalanceResponse](../../models/operations/createbalanceresponse.md), error**


## CreateWallet

Create a new wallet

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

    ctx := context.Background()
    res, err := s.Wallets.CreateWallet(ctx, &shared.CreateWalletRequest{
        Metadata: map[string]interface{}{
            "key": "string",
        },
        Name: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateWalletResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.CreateWalletRequest](../../models/shared/createwalletrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CreateWalletResponse](../../models/operations/createwalletresponse.md), error**


## CreditWallet

Credit a wallet

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


    var id string = "string"

    creditWalletRequest := &shared.CreditWalletRequest{
        Amount: shared.Monetary{
            Amount: 201874,
            Asset: "string",
        },
        Metadata: map[string]interface{}{
            "key": "string",
        },
        Sources: []shared.Subject{
            shared.CreateSubjectLedgerAccountSubject(
                shared.LedgerAccountSubject{
                    Identifier: "string",
                    Type: "string",
                },
            ),
        },
    }

    ctx := context.Background()
    res, err := s.Wallets.CreditWallet(ctx, id, creditWalletRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `id`                                                                      | *string*                                                                  | :heavy_check_mark:                                                        | N/A                                                                       |
| `creditWalletRequest`                                                     | [*shared.CreditWalletRequest](../../models/shared/creditwalletrequest.md) | :heavy_minus_sign:                                                        | N/A                                                                       |


### Response

**[*operations.CreditWalletResponse](../../models/operations/creditwalletresponse.md), error**


## DebitWallet

Debit a wallet

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


    var id string = "string"

    debitWalletRequest := &shared.DebitWalletRequest{
        Amount: shared.Monetary{
            Amount: 245256,
            Asset: "string",
        },
        Balances: []string{
            "string",
        },
        Destination: shared.CreateSubjectLedgerAccountSubject(
                shared.LedgerAccountSubject{
                    Identifier: "string",
                    Type: "string",
                },
        ),
        Metadata: map[string]interface{}{
            "key": "string",
        },
    }

    ctx := context.Background()
    res, err := s.Wallets.DebitWallet(ctx, id, debitWalletRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.DebitWalletResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |
| `id`                                                                    | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `debitWalletRequest`                                                    | [*shared.DebitWalletRequest](../../models/shared/debitwalletrequest.md) | :heavy_minus_sign:                                                      | N/A                                                                     |


### Response

**[*operations.DebitWalletResponse](../../models/operations/debitwalletresponse.md), error**


## GetBalance

Get detailed balance

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


    var balanceName string = "string"

    var id string = "string"

    ctx := context.Background()
    res, err := s.Wallets.GetBalance(ctx, balanceName, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetBalanceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `balanceName`                                         | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetBalanceResponse](../../models/operations/getbalanceresponse.md), error**


## GetHold

Get a hold

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


    var holdID string = "string"

    ctx := context.Background()
    res, err := s.Wallets.GetHold(ctx, holdID)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetHoldResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `holdID`                                              | *string*                                              | :heavy_check_mark:                                    | The hold ID                                           |


### Response

**[*operations.GetHoldResponse](../../models/operations/getholdresponse.md), error**


## GetHolds

Get all holds for a wallet

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


    var cursor *string = "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="

    metadata := &operations.GetHoldsMetadata{}

    var pageSize *int64 = 692494

    var walletID *string = "string"

    ctx := context.Background()
    res, err := s.Wallets.GetHolds(ctx, cursor, metadata, pageSize, walletID)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetHoldsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                           | Required                                                                                                                                                                                                                       | Description                                                                                                                                                                                                                    | Example                                                                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                                            |                                                                                                                                                                                                                                |
| `cursor`                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                             | Parameter used in pagination requests.<br/>Set to the value of next for the next page of results.<br/>Set to the value of previous for the previous page of results.<br/>No other parameters can be set when the pagination token is set.<br/> | aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==                                                                                                                                                                                   |
| `metadata`                                                                                                                                                                                                                     | [*operations.GetHoldsMetadata](../../models/operations/getholdsmetadata.md)                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                             | Filter holds by metadata key value pairs. Nested objects can be used as seen in the example below.                                                                                                                             |                                                                                                                                                                                                                                |
| `pageSize`                                                                                                                                                                                                                     | **int64*                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                             | The maximum number of results to return per page                                                                                                                                                                               |                                                                                                                                                                                                                                |
| `walletID`                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                             | The wallet to filter on                                                                                                                                                                                                        |                                                                                                                                                                                                                                |


### Response

**[*operations.GetHoldsResponse](../../models/operations/getholdsresponse.md), error**


## GetTransactions

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


    var cursor *string = "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="

    var pageSize *int64 = 680555

    var walletID *string = "string"

    ctx := context.Background()
    res, err := s.Wallets.GetTransactions(ctx, cursor, pageSize, walletID)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                            | Type                                                                                                                                                                                                                 | Required                                                                                                                                                                                                             | Description                                                                                                                                                                                                          | Example                                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                  |                                                                                                                                                                                                                      |
| `cursor`                                                                                                                                                                                                             | **string*                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                   | Parameter used in pagination requests.<br/>Set to the value of next for the next page of results.<br/>Set to the value of previous for the previous page of results.<br/>No other parameters can be set when the cursor is set.<br/> | aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==                                                                                                                                                                         |
| `pageSize`                                                                                                                                                                                                           | **int64*                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                   | The maximum number of results to return per page                                                                                                                                                                     |                                                                                                                                                                                                                      |
| `walletID`                                                                                                                                                                                                           | **string*                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                   | A wallet ID to filter on                                                                                                                                                                                             |                                                                                                                                                                                                                      |


### Response

**[*operations.GetTransactionsResponse](../../models/operations/gettransactionsresponse.md), error**


## GetWallet

Get a wallet

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


    var id string = "string"

    ctx := context.Background()
    res, err := s.Wallets.GetWallet(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetWalletResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetWalletResponse](../../models/operations/getwalletresponse.md), error**


## ListBalances

List balances of a wallet

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


    var id string = "string"

    ctx := context.Background()
    res, err := s.Wallets.ListBalances(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListBalancesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.ListBalancesResponse](../../models/operations/listbalancesresponse.md), error**


## ListWallets

List all wallets

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


    var cursor *string = "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="

    metadata := &operations.ListWalletsMetadata{}

    var name *string = "string"

    var pageSize *int64 = 412533

    ctx := context.Background()
    res, err := s.Wallets.ListWallets(ctx, cursor, metadata, name, pageSize)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListWalletsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                           | Required                                                                                                                                                                                                                       | Description                                                                                                                                                                                                                    | Example                                                                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                                            |                                                                                                                                                                                                                                |
| `cursor`                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                             | Parameter used in pagination requests.<br/>Set to the value of next for the next page of results.<br/>Set to the value of previous for the previous page of results.<br/>No other parameters can be set when the pagination token is set.<br/> | aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==                                                                                                                                                                                   |
| `metadata`                                                                                                                                                                                                                     | [*operations.ListWalletsMetadata](../../models/operations/listwalletsmetadata.md)                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                             | Filter wallets by metadata key value pairs. Nested objects can be used as seen in the example below.                                                                                                                           |                                                                                                                                                                                                                                |
| `name`                                                                                                                                                                                                                         | **string*                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                             | Filter on wallet name                                                                                                                                                                                                          |                                                                                                                                                                                                                                |
| `pageSize`                                                                                                                                                                                                                     | **int64*                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                             | The maximum number of results to return per page                                                                                                                                                                               |                                                                                                                                                                                                                                |


### Response

**[*operations.ListWalletsResponse](../../models/operations/listwalletsresponse.md), error**


## UpdateWallet

Update a wallet

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


    var id string = "string"

    requestBody := &operations.UpdateWalletRequestBody{
        Metadata: map[string]interface{}{
            "key": "string",
        },
    }

    ctx := context.Background()
    res, err := s.Wallets.UpdateWallet(ctx, id, requestBody)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |
| `id`                                                                                      | *string*                                                                                  | :heavy_check_mark:                                                                        | N/A                                                                                       |
| `requestBody`                                                                             | [*operations.UpdateWalletRequestBody](../../models/operations/updatewalletrequestbody.md) | :heavy_minus_sign:                                                                        | N/A                                                                                       |


### Response

**[*operations.UpdateWalletResponse](../../models/operations/updatewalletresponse.md), error**


## VoidHold

Cancel a hold

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


    var holdID string = "string"

    ctx := context.Background()
    res, err := s.Wallets.VoidHold(ctx, holdID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `holdID`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.VoidHoldResponse](../../models/operations/voidholdresponse.md), error**


## WalletsgetServerInfo

Get server info

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

    ctx := context.Background()
    res, err := s.Wallets.WalletsgetServerInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.WalletsgetServerInfoResponse](../../models/operations/walletsgetserverinforesponse.md), error**

