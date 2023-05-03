# Scopes

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
	"github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Scopes.AddTransientScope(ctx, operations.AddTransientScopeRequest{
        ScopeID: "accusamus",
        TransientScopeID: "commodi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CreateScope

Create scope

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Scopes.CreateScope(ctx, shared.CreateScopeRequest{
        Label: "repudiandae",
        Metadata: map[string]interface{}{
            "ipsum": "quidem",
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

## DeleteScope

Delete scope

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
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Scopes.DeleteScope(ctx, operations.DeleteScopeRequest{
        ScopeID: "molestias",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## DeleteTransientScope

Delete a transient scope from a scope

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
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Scopes.DeleteTransientScope(ctx, operations.DeleteTransientScopeRequest{
        ScopeID: "excepturi",
        TransientScopeID: "pariatur",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## ListScopes

List Scopes

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/formance-go-sdk"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
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

## ReadScope

Read scope

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
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Scopes.ReadScope(ctx, operations.ReadScopeRequest{
        ScopeID: "modi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReadScopeResponse != nil {
        // handle response
    }
}
```

## UpdateScope

Update scope

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Scopes.UpdateScope(ctx, operations.UpdateScopeRequest{
        UpdateScopeRequest: &shared.UpdateScopeRequest{
            Label: "praesentium",
            Metadata: map[string]interface{}{
                "voluptates": "quasi",
                "repudiandae": "sint",
                "veritatis": "itaque",
            },
        },
        ScopeID: "incidunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateScopeResponse != nil {
        // handle response
    }
}
```
