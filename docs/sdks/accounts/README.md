# Accounts
(*Accounts*)

## Overview

Everything related to Accounts

### Available Operations

* [AddMetadataToAccount](#addmetadatatoaccount) - Add metadata to an account
* [CountAccounts](#countaccounts) - Count the accounts from a ledger
* [GetAccount](#getaccount) - Get account by its address
* [ListAccounts](#listaccounts) - List accounts from a ledger

## AddMetadataToAccount

Add metadata to an account

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
        formancegosdk.WithSecurity(""),
    )


    requestBody := map[string]interface{}{
        "a": "string",
        "admin": "string",
    }

    var address string = "users:001"

    var ledger string = "ledger001"

    ctx := context.Background()
    res, err := s.Accounts.AddMetadataToAccount(ctx, requestBody, address, ledger)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |                                                                                                              |
| `requestBody`                                                                                                | map[string]*interface{}*                                                                                     | :heavy_check_mark:                                                                                           | metadata                                                                                                     | [object Object]                                                                                              |
| `address`                                                                                                    | *string*                                                                                                     | :heavy_check_mark:                                                                                           | Exact address of the account. It must match the following regular expressions pattern:<br/>```<br/>^\w+(:\w+)*$<br/>```<br/> | users:001                                                                                                    |
| `ledger`                                                                                                     | *string*                                                                                                     | :heavy_check_mark:                                                                                           | Name of the ledger.                                                                                          | ledger001                                                                                                    |


### Response

**[*operations.AddMetadataToAccountResponse](../../pkg/models/operations/addmetadatatoaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CountAccounts

Count the accounts from a ledger

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
        formancegosdk.WithSecurity(""),
    )


    var ledger string = "ledger001"

    var address *string = "users:.+"

    metadata := &operations.Metadata{}

    ctx := context.Background()
    res, err := s.Accounts.CountAccounts(ctx, ledger, address, metadata)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           | Example                                                                                               |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |                                                                                                       |
| `ledger`                                                                                              | *string*                                                                                              | :heavy_check_mark:                                                                                    | Name of the ledger.                                                                                   | ledger001                                                                                             |
| `address`                                                                                             | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter accounts by address pattern (regular expression placed between ^ and $).                       | users:.+                                                                                              |
| `metadata`                                                                                            | [*operations.Metadata](../../../pkg/models/operations/metadata.md)                                    | :heavy_minus_sign:                                                                                    | Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below. |                                                                                                       |


### Response

**[*operations.CountAccountsResponse](../../pkg/models/operations/countaccountsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetAccount

Get account by its address

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
        formancegosdk.WithSecurity(""),
    )


    var address string = "users:001"

    var ledger string = "ledger001"

    ctx := context.Background()
    res, err := s.Accounts.GetAccount(ctx, address, ledger)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |                                                                                                              |
| `address`                                                                                                    | *string*                                                                                                     | :heavy_check_mark:                                                                                           | Exact address of the account. It must match the following regular expressions pattern:<br/>```<br/>^\w+(:\w+)*$<br/>```<br/> | users:001                                                                                                    |
| `ledger`                                                                                                     | *string*                                                                                                     | :heavy_check_mark:                                                                                           | Name of the ledger.                                                                                          | ledger001                                                                                                    |


### Response

**[*operations.GetAccountResponse](../../pkg/models/operations/getaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListAccounts

List accounts from a ledger, sorted by address in descending order.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"context"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Accounts.ListAccounts(ctx, operations.ListAccountsRequest{
        Address: formancegosdk.String("users:.+"),
        After: formancegosdk.String("users:003"),
        Balance: formancegosdk.Int64(2400),
        BalanceOperator: operations.BalanceOperatorGte.ToPointer(),
        BalanceOperatorDeprecated: operations.QueryParamBalanceOperatorGte.ToPointer(),
        Cursor: formancegosdk.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Ledger: "ledger001",
        Metadata: &operations.QueryParamMetadata{},
        PaginationToken: formancegosdk.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListAccountsRequest](../../pkg/models/operations/listaccountsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListAccountsResponse](../../pkg/models/operations/listaccountsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
