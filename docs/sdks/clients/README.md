# Clients
(*Clients*)

## Overview

Everything related to Clients

### Available Operations

* [AddScopeToClient](#addscopetoclient) - Add scope to client
* [CreateClient](#createclient) - Create client
* [CreateSecret](#createsecret) - Add a secret to a client
* [DeleteClient](#deleteclient) - Delete client
* [DeleteScopeFromClient](#deletescopefromclient) - Delete scope from client
* [DeleteSecret](#deletesecret) - Delete a secret from a client
* [ListClients](#listclients) - List clients
* [ReadClient](#readclient) - Read client
* [UpdateClient](#updateclient) - Update client

## AddScopeToClient

Add scope to client

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


    var clientID string = "<value>"

    var scopeID string = "<value>"

    ctx := context.Background()
    res, err := s.Clients.AddScopeToClient(ctx, clientID, scopeID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `clientID`                                            | *string*                                              | :heavy_check_mark:                                    | Client ID                                             |
| `scopeID`                                             | *string*                                              | :heavy_check_mark:                                    | Scope ID                                              |


### Response

**[*operations.AddScopeToClientResponse](../../pkg/models/operations/addscopetoclientresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateClient

Create client

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
    res, err := s.Clients.CreateClient(ctx, &shared.CreateClientRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateClientResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.CreateClientRequest](../../pkg/models/shared/createclientrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.CreateClientResponse](../../pkg/models/operations/createclientresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateSecret

Add a secret to a client

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


    var clientID string = "<value>"

    createSecretRequest := &shared.CreateSecretRequest{
        Name: "<value>",
    }

    ctx := context.Background()
    res, err := s.Clients.CreateSecret(ctx, clientID, createSecretRequest)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateSecretResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                     | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ctx`                                                                         | [context.Context](https://pkg.go.dev/context#Context)                         | :heavy_check_mark:                                                            | The context to use for the request.                                           |
| `clientID`                                                                    | *string*                                                                      | :heavy_check_mark:                                                            | Client ID                                                                     |
| `createSecretRequest`                                                         | [*shared.CreateSecretRequest](../../pkg/models/shared/createsecretrequest.md) | :heavy_minus_sign:                                                            | N/A                                                                           |


### Response

**[*operations.CreateSecretResponse](../../pkg/models/operations/createsecretresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteClient

Delete client

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


    var clientID string = "<value>"

    ctx := context.Background()
    res, err := s.Clients.DeleteClient(ctx, clientID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `clientID`                                            | *string*                                              | :heavy_check_mark:                                    | Client ID                                             |


### Response

**[*operations.DeleteClientResponse](../../pkg/models/operations/deleteclientresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteScopeFromClient

Delete scope from client

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


    var clientID string = "<value>"

    var scopeID string = "<value>"

    ctx := context.Background()
    res, err := s.Clients.DeleteScopeFromClient(ctx, clientID, scopeID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `clientID`                                            | *string*                                              | :heavy_check_mark:                                    | Client ID                                             |
| `scopeID`                                             | *string*                                              | :heavy_check_mark:                                    | Scope ID                                              |


### Response

**[*operations.DeleteScopeFromClientResponse](../../pkg/models/operations/deletescopefromclientresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteSecret

Delete a secret from a client

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


    var clientID string = "<value>"

    var secretID string = "<value>"

    ctx := context.Background()
    res, err := s.Clients.DeleteSecret(ctx, clientID, secretID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `clientID`                                            | *string*                                              | :heavy_check_mark:                                    | Client ID                                             |
| `secretID`                                            | *string*                                              | :heavy_check_mark:                                    | Secret ID                                             |


### Response

**[*operations.DeleteSecretResponse](../../pkg/models/operations/deletesecretresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListClients

List clients

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
    res, err := s.Clients.ListClients(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListClientsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListClientsResponse](../../pkg/models/operations/listclientsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReadClient

Read client

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


    var clientID string = "<value>"

    ctx := context.Background()
    res, err := s.Clients.ReadClient(ctx, clientID)
    if err != nil {
        log.Fatal(err)
    }
    if res.ReadClientResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `clientID`                                            | *string*                                              | :heavy_check_mark:                                    | Client ID                                             |


### Response

**[*operations.ReadClientResponse](../../pkg/models/operations/readclientresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateClient

Update client

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


    var clientID string = "<value>"

    updateClientRequest := &shared.UpdateClientRequest{
        Name: "<value>",
    }

    ctx := context.Background()
    res, err := s.Clients.UpdateClient(ctx, clientID, updateClientRequest)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateClientResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                     | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ctx`                                                                         | [context.Context](https://pkg.go.dev/context#Context)                         | :heavy_check_mark:                                                            | The context to use for the request.                                           |
| `clientID`                                                                    | *string*                                                                      | :heavy_check_mark:                                                            | Client ID                                                                     |
| `updateClientRequest`                                                         | [*shared.UpdateClientRequest](../../pkg/models/shared/updateclientrequest.md) | :heavy_minus_sign:                                                            | N/A                                                                           |


### Response

**[*operations.UpdateClientResponse](../../pkg/models/operations/updateclientresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
