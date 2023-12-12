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
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
	"net/http"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ctx`                                                                       | [context.Context](https://pkg.go.dev/context#Context)                       | :heavy_check_mark:                                                          | The context to use for the request.                                         |
| `holdID`                                                                    | *string*                                                                    | :heavy_check_mark:                                                          | N/A                                                                         |
| `confirmHoldRequest`                                                        | [*shared.ConfirmHoldRequest](../../pkg/models/shared/confirmholdrequest.md) | :heavy_minus_sign:                                                          | N/A                                                                         |


### Response

**[*operations.ConfirmHoldResponse](../../pkg/models/operations/confirmholdresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateBalance

Create a balance

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ctx`                                                                           | [context.Context](https://pkg.go.dev/context#Context)                           | :heavy_check_mark:                                                              | The context to use for the request.                                             |
| `id`                                                                            | *string*                                                                        | :heavy_check_mark:                                                              | N/A                                                                             |
| `createBalanceRequest`                                                          | [*shared.CreateBalanceRequest](../../pkg/models/shared/createbalancerequest.md) | :heavy_minus_sign:                                                              | N/A                                                                             |


### Response

**[*operations.CreateBalanceResponse](../../pkg/models/operations/createbalanceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateWallet

Create a new wallet

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.CreateWalletRequest](../../pkg/models/shared/createwalletrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.CreateWalletResponse](../../pkg/models/operations/createwalletresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreditWallet

Credit a wallet

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
	"net/http"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

| Parameter                                                                     | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ctx`                                                                         | [context.Context](https://pkg.go.dev/context#Context)                         | :heavy_check_mark:                                                            | The context to use for the request.                                           |
| `id`                                                                          | *string*                                                                      | :heavy_check_mark:                                                            | N/A                                                                           |
| `creditWalletRequest`                                                         | [*shared.CreditWalletRequest](../../pkg/models/shared/creditwalletrequest.md) | :heavy_minus_sign:                                                            | N/A                                                                           |


### Response

**[*operations.CreditWalletResponse](../../pkg/models/operations/creditwalletresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DebitWallet

Debit a wallet

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ctx`                                                                       | [context.Context](https://pkg.go.dev/context#Context)                       | :heavy_check_mark:                                                          | The context to use for the request.                                         |
| `id`                                                                        | *string*                                                                    | :heavy_check_mark:                                                          | N/A                                                                         |
| `debitWalletRequest`                                                        | [*shared.DebitWalletRequest](../../pkg/models/shared/debitwalletrequest.md) | :heavy_minus_sign:                                                          | N/A                                                                         |


### Response

**[*operations.DebitWalletResponse](../../pkg/models/operations/debitwalletresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetBalance

Get detailed balance

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

**[*operations.GetBalanceResponse](../../pkg/models/operations/getbalanceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetHold

Get a hold

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

**[*operations.GetHoldResponse](../../pkg/models/operations/getholdresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetHolds

Get all holds for a wallet

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )


    var cursor *string = "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="

    metadata := &operations.GetHoldsQueryParamMetadata{}

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
| `metadata`                                                                                                                                                                                                                     | [*operations.GetHoldsQueryParamMetadata](../../pkg/models/operations/getholdsqueryparammetadata.md)                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                             | Filter holds by metadata key value pairs. Nested objects can be used as seen in the example below.                                                                                                                             |                                                                                                                                                                                                                                |
| `pageSize`                                                                                                                                                                                                                     | **int64*                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                             | The maximum number of results to return per page                                                                                                                                                                               |                                                                                                                                                                                                                                |
| `walletID`                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                             | The wallet to filter on                                                                                                                                                                                                        |                                                                                                                                                                                                                                |


### Response

**[*operations.GetHoldsResponse](../../pkg/models/operations/getholdsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetTransactions

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

**[*operations.GetTransactionsResponse](../../pkg/models/operations/gettransactionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetWallet

Get a wallet

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

**[*operations.GetWalletResponse](../../pkg/models/operations/getwalletresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListBalances

List balances of a wallet

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

**[*operations.ListBalancesResponse](../../pkg/models/operations/listbalancesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListWallets

List all wallets

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )


    var cursor *string = "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="

    metadata := &operations.ListWalletsQueryParamMetadata{}

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
| `metadata`                                                                                                                                                                                                                     | [*operations.ListWalletsQueryParamMetadata](../../pkg/models/operations/listwalletsqueryparammetadata.md)                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                             | Filter wallets by metadata key value pairs. Nested objects can be used as seen in the example below.                                                                                                                           |                                                                                                                                                                                                                                |
| `name`                                                                                                                                                                                                                         | **string*                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                             | Filter on wallet name                                                                                                                                                                                                          |                                                                                                                                                                                                                                |
| `pageSize`                                                                                                                                                                                                                     | **int64*                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                             | The maximum number of results to return per page                                                                                                                                                                               |                                                                                                                                                                                                                                |


### Response

**[*operations.ListWalletsResponse](../../pkg/models/operations/listwalletsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateWallet

Update a wallet

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
	"context"
	"log"
	"net/http"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |
| `id`                                                                                          | *string*                                                                                      | :heavy_check_mark:                                                                            | N/A                                                                                           |
| `requestBody`                                                                                 | [*operations.UpdateWalletRequestBody](../../pkg/models/operations/updatewalletrequestbody.md) | :heavy_minus_sign:                                                                            | N/A                                                                                           |


### Response

**[*operations.UpdateWalletResponse](../../pkg/models/operations/updatewalletresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## VoidHold

Cancel a hold

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
	"net/http"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

**[*operations.VoidHoldResponse](../../pkg/models/operations/voidholdresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## WalletsgetServerInfo

Get server info

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

**[*operations.WalletsgetServerInfoResponse](../../pkg/models/operations/walletsgetserverinforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
