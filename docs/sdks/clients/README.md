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
	"context"
	"log"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
    )


    var clientID string = "string"

    var scopeID string = "string"

    ctx := context.Background()
    res, err := s.Clients.AddScopeToClient(ctx, clientID, scopeID)
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
| `clientID`                                            | *string*                                              | :heavy_check_mark:                                    | Client ID                                             |
| `scopeID`                                             | *string*                                              | :heavy_check_mark:                                    | Scope ID                                              |


### Response

**[*operations.AddScopeToClientResponse](../../models/operations/addscopetoclientresponse.md), error**


## CreateClient

Create client

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
    res, err := s.Clients.CreateClient(ctx, &shared.CreateClientRequest{
        Metadata: map[string]interface{}{
            "key": "string",
        },
        Name: "string",
        PostLogoutRedirectUris: []string{
            "string",
        },
        RedirectUris: []string{
            "string",
        },
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.CreateClientRequest](../../models/shared/createclientrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CreateClientResponse](../../models/operations/createclientresponse.md), error**


## CreateSecret

Add a secret to a client

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


    var clientID string = "string"

    createSecretRequest := &shared.CreateSecretRequest{
        Metadata: map[string]interface{}{
            "key": "string",
        },
        Name: "string",
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

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `clientID`                                                                | *string*                                                                  | :heavy_check_mark:                                                        | Client ID                                                                 |
| `createSecretRequest`                                                     | [*shared.CreateSecretRequest](../../models/shared/createsecretrequest.md) | :heavy_minus_sign:                                                        | N/A                                                                       |


### Response

**[*operations.CreateSecretResponse](../../models/operations/createsecretresponse.md), error**


## DeleteClient

Delete client

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


    var clientID string = "string"

    ctx := context.Background()
    res, err := s.Clients.DeleteClient(ctx, clientID)
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
| `clientID`                                            | *string*                                              | :heavy_check_mark:                                    | Client ID                                             |


### Response

**[*operations.DeleteClientResponse](../../models/operations/deleteclientresponse.md), error**


## DeleteScopeFromClient

Delete scope from client

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


    var clientID string = "string"

    var scopeID string = "string"

    ctx := context.Background()
    res, err := s.Clients.DeleteScopeFromClient(ctx, clientID, scopeID)
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
| `clientID`                                            | *string*                                              | :heavy_check_mark:                                    | Client ID                                             |
| `scopeID`                                             | *string*                                              | :heavy_check_mark:                                    | Scope ID                                              |


### Response

**[*operations.DeleteScopeFromClientResponse](../../models/operations/deletescopefromclientresponse.md), error**


## DeleteSecret

Delete a secret from a client

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


    var clientID string = "string"

    var secretID string = "string"

    ctx := context.Background()
    res, err := s.Clients.DeleteSecret(ctx, clientID, secretID)
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
| `clientID`                                            | *string*                                              | :heavy_check_mark:                                    | Client ID                                             |
| `secretID`                                            | *string*                                              | :heavy_check_mark:                                    | Secret ID                                             |


### Response

**[*operations.DeleteSecretResponse](../../models/operations/deletesecretresponse.md), error**


## ListClients

List clients

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

**[*operations.ListClientsResponse](../../models/operations/listclientsresponse.md), error**


## ReadClient

Read client

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


    var clientID string = "string"

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

**[*operations.ReadClientResponse](../../models/operations/readclientresponse.md), error**


## UpdateClient

Update client

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


    var clientID string = "string"

    updateClientRequest := &shared.UpdateClientRequest{
        Metadata: map[string]interface{}{
            "key": "string",
        },
        Name: "string",
        PostLogoutRedirectUris: []string{
            "string",
        },
        RedirectUris: []string{
            "string",
        },
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

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `clientID`                                                                | *string*                                                                  | :heavy_check_mark:                                                        | Client ID                                                                 |
| `updateClientRequest`                                                     | [*shared.UpdateClientRequest](../../models/shared/updateclientrequest.md) | :heavy_minus_sign:                                                        | N/A                                                                       |


### Response

**[*operations.UpdateClientResponse](../../models/operations/updateclientresponse.md), error**

