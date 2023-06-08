# Users

## Overview

Everything related to Users

### Available Operations

* [ListUsers](#listusers) - List users
* [ReadUser](#readuser) - Read user

## ListUsers

List users

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
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Users.ListUsers(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListUsersResponse != nil {
        // handle response
    }
}
```

## ReadUser

Read user

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
    res, err := s.Users.ReadUser(ctx, operations.ReadUserRequest{
        UserID: "excepturi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReadUserResponse != nil {
        // handle response
    }
}
```
