# Payments
(*Payments*)

## Overview

Everything related to Payments

### Available Operations

* [ConnectorsStripeTransfer](#connectorsstripetransfer) - Transfer funds between Stripe accounts
* [GetConnectorTask](#getconnectortask) - Read a specific task of the connector
* [GetPayment](#getpayment) - Get a payment
* [InstallConnector](#installconnector) - Install a connector
* [ListAllConnectors](#listallconnectors) - List all installed connectors
* [ListConfigsAvailableConnectors](#listconfigsavailableconnectors) - List the configs of each available connector
* [ListConnectorTasks](#listconnectortasks) - List tasks from a connector
* [ListPayments](#listpayments) - List payments
* [PaymentslistAccounts](#paymentslistaccounts) - List accounts
* [ReadConnectorConfig](#readconnectorconfig) - Read the config of a connector
* [ResetConnector](#resetconnector) - Reset a connector
* [UninstallConnector](#uninstallconnector) - Uninstall a connector

## ConnectorsStripeTransfer

Execute a transfer between two Stripe accounts.

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
    res, err := s.Payments.ConnectorsStripeTransfer(ctx, shared.StripeTransferRequest{
        Amount: formancegosdk.Int64(100),
        Asset: formancegosdk.String("USD"),
        Destination: formancegosdk.String("acct_1Gqj58KZcSIg2N2q"),
        Metadata: &shared.StripeTransferRequestMetadata{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StripeTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.StripeTransferRequest](../../models/shared/stripetransferrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.ConnectorsStripeTransferResponse](../../models/operations/connectorsstripetransferresponse.md), error**


## GetConnectorTask

Get a specific task associated to the connector.

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


    var connector shared.Connector = shared.ConnectorBankingCircle

    var taskID string = "string"

    ctx := context.Background()
    res, err := s.Payments.GetConnectorTask(ctx, connector, taskID)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `connector`                                           | [shared.Connector](../../models/shared/connector.md)  | :heavy_check_mark:                                    | The name of the connector.                            |
| `taskID`                                              | *string*                                              | :heavy_check_mark:                                    | The task ID.                                          |


### Response

**[*operations.GetConnectorTaskResponse](../../models/operations/getconnectortaskresponse.md), error**


## GetPayment

Get a payment

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


    var paymentID string = "string"

    ctx := context.Background()
    res, err := s.Payments.GetPayment(ctx, paymentID)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `paymentID`                                           | *string*                                              | :heavy_check_mark:                                    | The payment ID.                                       |


### Response

**[*operations.GetPaymentResponse](../../models/operations/getpaymentresponse.md), error**


## InstallConnector

Install a connector by its name and config.

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


    var connectorConfig shared.ConnectorConfig = shared.CreateConnectorConfigDummyPayConfig(
            shared.DummyPayConfig{
                Directory: "/tmp/dummypay",
                FileGenerationPeriod: formancegosdk.String("60s"),
                FilePollingPeriod: formancegosdk.String("60s"),
            },
    )

    var connector shared.Connector = shared.ConnectorBankingCircle

    ctx := context.Background()
    res, err := s.Payments.InstallConnector(ctx, connectorConfig, connector)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `connectorConfig`                                                | [shared.ConnectorConfig](../../models/shared/connectorconfig.md) | :heavy_check_mark:                                               | N/A                                                              |
| `connector`                                                      | [shared.Connector](../../models/shared/connector.md)             | :heavy_check_mark:                                               | The name of the connector.                                       |


### Response

**[*operations.InstallConnectorResponse](../../models/operations/installconnectorresponse.md), error**


## ListAllConnectors

List all installed connectors.

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
    res, err := s.Payments.ListAllConnectors(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListAllConnectorsResponse](../../models/operations/listallconnectorsresponse.md), error**


## ListConfigsAvailableConnectors

List the configs of each available connector.

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
    res, err := s.Payments.ListConfigsAvailableConnectors(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorsConfigsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListConfigsAvailableConnectorsResponse](../../models/operations/listconfigsavailableconnectorsresponse.md), error**


## ListConnectorTasks

List all tasks associated with this connector.

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


    var connector shared.Connector = shared.ConnectorDummyPay

    var cursor *string = "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="

    var pageSize *int64 = 501686

    ctx := context.Background()
    res, err := s.Payments.ListConnectorTasks(ctx, connector, cursor, pageSize)
    if err != nil {
        log.Fatal(err)
    }

    if res.TasksCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                          |
| `connector`                                                                                                                                                                                                                                              | [shared.Connector](../../models/shared/connector.md)                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                       | The name of the connector.                                                                                                                                                                                                                               |                                                                                                                                                                                                                                                          |
| `cursor`                                                                                                                                                                                                                                                 | **string*                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                       | Parameter used in pagination requests. Maximum page size is set to 15.<br/>Set to the value of next for the next page of results.<br/>Set to the value of previous for the previous page of results.<br/>No other parameters can be set when this parameter is set.<br/> | aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==                                                                                                                                                                                                             |
| `pageSize`                                                                                                                                                                                                                                               | **int64*                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                       | The maximum number of results to return per page.<br/>                                                                                                                                                                                                   |                                                                                                                                                                                                                                                          |


### Response

**[*operations.ListConnectorTasksResponse](../../models/operations/listconnectortasksresponse.md), error**


## ListPayments

List payments

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

    var pageSize *int64 = 13778

    sort := []string{
        "string",
    }

    ctx := context.Background()
    res, err := s.Payments.ListPayments(ctx, cursor, pageSize, sort)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentsCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                          |
| `cursor`                                                                                                                                                                                                                                                 | **string*                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                       | Parameter used in pagination requests. Maximum page size is set to 15.<br/>Set to the value of next for the next page of results.<br/>Set to the value of previous for the previous page of results.<br/>No other parameters can be set when this parameter is set.<br/> | aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==                                                                                                                                                                                                             |
| `pageSize`                                                                                                                                                                                                                                               | **int64*                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                       | The maximum number of results to return per page.<br/>                                                                                                                                                                                                   |                                                                                                                                                                                                                                                          |
| `sort`                                                                                                                                                                                                                                                   | []*string*                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                       | Fields used to sort payments (default is date:desc).                                                                                                                                                                                                     |                                                                                                                                                                                                                                                          |


### Response

**[*operations.ListPaymentsResponse](../../models/operations/listpaymentsresponse.md), error**


## PaymentslistAccounts

List accounts

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

    var pageSize *int64 = 227071

    sort := []string{
        "string",
    }

    ctx := context.Background()
    res, err := s.Payments.PaymentslistAccounts(ctx, cursor, pageSize, sort)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountsCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                          |
| `cursor`                                                                                                                                                                                                                                                 | **string*                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                       | Parameter used in pagination requests. Maximum page size is set to 15.<br/>Set to the value of next for the next page of results.<br/>Set to the value of previous for the previous page of results.<br/>No other parameters can be set when this parameter is set.<br/> | aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==                                                                                                                                                                                                             |
| `pageSize`                                                                                                                                                                                                                                               | **int64*                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                       | The maximum number of results to return per page.<br/>                                                                                                                                                                                                   |                                                                                                                                                                                                                                                          |
| `sort`                                                                                                                                                                                                                                                   | []*string*                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                       | Fields used to sort payments (default is date:desc).                                                                                                                                                                                                     |                                                                                                                                                                                                                                                          |


### Response

**[*operations.PaymentslistAccountsResponse](../../models/operations/paymentslistaccountsresponse.md), error**


## ReadConnectorConfig

Read connector config

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


    var connector shared.Connector = shared.ConnectorBankingCircle

    ctx := context.Background()
    res, err := s.Payments.ReadConnectorConfig(ctx, connector)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectorConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `connector`                                           | [shared.Connector](../../models/shared/connector.md)  | :heavy_check_mark:                                    | The name of the connector.                            |


### Response

**[*operations.ReadConnectorConfigResponse](../../models/operations/readconnectorconfigresponse.md), error**


## ResetConnector

Reset a connector by its name.
It will remove the connector and ALL PAYMENTS generated with it.


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


    var connector shared.Connector = shared.ConnectorCurrencyCloud

    ctx := context.Background()
    res, err := s.Payments.ResetConnector(ctx, connector)
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
| `connector`                                           | [shared.Connector](../../models/shared/connector.md)  | :heavy_check_mark:                                    | The name of the connector.                            |


### Response

**[*operations.ResetConnectorResponse](../../models/operations/resetconnectorresponse.md), error**


## UninstallConnector

Uninstall a connector by its name.

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


    var connector shared.Connector = shared.ConnectorDummyPay

    ctx := context.Background()
    res, err := s.Payments.UninstallConnector(ctx, connector)
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
| `connector`                                           | [shared.Connector](../../models/shared/connector.md)  | :heavy_check_mark:                                    | The name of the connector.                            |


### Response

**[*operations.UninstallConnectorResponse](../../models/operations/uninstallconnectorresponse.md), error**

