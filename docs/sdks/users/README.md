# Users
(*Users*)

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
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

func main() {
    s := formancegosdk.New(
        formancegosdk.WithSecurity(""),
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListUsersResponse](../../models/operations/listusersresponse.md), error**


## ReadUser

Read user

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


    var userID string = "string"

    ctx := context.Background()
    res, err := s.Users.ReadUser(ctx, userID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ReadUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `userID`                                              | *string*                                              | :heavy_check_mark:                                    | User ID                                               |


### Response

**[*operations.ReadUserResponse](../../models/operations/readuserresponse.md), error**

