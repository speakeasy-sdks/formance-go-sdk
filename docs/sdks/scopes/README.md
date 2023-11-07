# Scopes
(*.Scopes*)

## Overview

Everything related to Scopes

### Available Operations

* [AddTransientScope](#addtransientscope) - Add a transient scope to a scope
* [CreateScope](#createscope) - Create scope
* [DeleteScope](#deletescope) - Delete scope
* [DeleteTransientScope](#deletetransientscope) - Delete a transient scope from a scope
* [ListScopes](#listscopes) - List scopes
* [ReadScope](#readscope) - Read scope
* [UpdateScope](#updatescope) - Update scope

## AddTransientScope

Add a transient scope to a scope

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


    var scopeID string = "string"

    var transientScopeID string = "string"

    ctx := context.Background()
    res, err := s.Scopes.AddTransientScope(ctx, scopeID, transientScopeID)
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
| `scopeID`                                             | *string*                                              | :heavy_check_mark:                                    | Scope ID                                              |
| `transientScopeID`                                    | *string*                                              | :heavy_check_mark:                                    | Transient scope ID                                    |


### Response

**[*operations.AddTransientScopeResponse](../../models/operations/addtransientscoperesponse.md), error**


## CreateScope

Create scope

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
    res, err := s.Scopes.CreateScope(ctx, &shared.CreateScopeRequest{
        Label: "string",
        Metadata: map[string]interface{}{
            "key": "string",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateScopeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.CreateScopeRequest](../../models/shared/createscoperequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.CreateScopeResponse](../../models/operations/createscoperesponse.md), error**


## DeleteScope

Delete scope

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


    var scopeID string = "string"

    ctx := context.Background()
    res, err := s.Scopes.DeleteScope(ctx, scopeID)
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
| `scopeID`                                             | *string*                                              | :heavy_check_mark:                                    | Scope ID                                              |


### Response

**[*operations.DeleteScopeResponse](../../models/operations/deletescoperesponse.md), error**


## DeleteTransientScope

Delete a transient scope from a scope

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


    var scopeID string = "string"

    var transientScopeID string = "string"

    ctx := context.Background()
    res, err := s.Scopes.DeleteTransientScope(ctx, scopeID, transientScopeID)
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
| `scopeID`                                             | *string*                                              | :heavy_check_mark:                                    | Scope ID                                              |
| `transientScopeID`                                    | *string*                                              | :heavy_check_mark:                                    | Transient scope ID                                    |


### Response

**[*operations.DeleteTransientScopeResponse](../../models/operations/deletetransientscoperesponse.md), error**


## ListScopes

List Scopes

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
    res, err := s.Scopes.ListScopes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListScopesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListScopesResponse](../../models/operations/listscopesresponse.md), error**


## ReadScope

Read scope

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


    var scopeID string = "string"

    ctx := context.Background()
    res, err := s.Scopes.ReadScope(ctx, scopeID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ReadScopeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `scopeID`                                             | *string*                                              | :heavy_check_mark:                                    | Scope ID                                              |


### Response

**[*operations.ReadScopeResponse](../../models/operations/readscoperesponse.md), error**


## UpdateScope

Update scope

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


    var scopeID string = "string"

    updateScopeRequest := &shared.UpdateScopeRequest{
        Label: "string",
        Metadata: map[string]interface{}{
            "key": "string",
        },
    }

    ctx := context.Background()
    res, err := s.Scopes.UpdateScope(ctx, scopeID, updateScopeRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateScopeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |
| `scopeID`                                                               | *string*                                                                | :heavy_check_mark:                                                      | Scope ID                                                                |
| `updateScopeRequest`                                                    | [*shared.UpdateScopeRequest](../../models/shared/updatescoperequest.md) | :heavy_minus_sign:                                                      | N/A                                                                     |


### Response

**[*operations.UpdateScopeResponse](../../models/operations/updatescoperesponse.md), error**

