# Script
(*Script*)

## Overview

Everything related to Script

### Available Operations

* [~~RunScript~~](#runscript) - Execute a Numscript :warning: **Deprecated**

## ~~RunScript~~

This route is deprecated, and has been merged into `POST /{ledger}/transactions`.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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


    script := shared.Script{
        Metadata: map[string]interface{}{
            "admin": "string",
            "a": "string",
        },
        Plain: "vars {
    account $user
    }
    send [COIN 10] (
    	source = @world
    	destination = $user
    )
    ",
        Reference: formancegosdk.String("order_1234"),
        Vars: &shared.Vars{},
    }

    var ledger string = "ledger001"

    var preview *bool = true

    ctx := context.Background()
    res, err := s.Script.RunScript(ctx, script, ledger, preview)
    if err != nil {
        log.Fatal(err)
    }

    if res.ScriptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                           | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         | Example                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                               | :heavy_check_mark:                                                                                                  | The context to use for the request.                                                                                 |                                                                                                                     |
| `script`                                                                                                            | [shared.Script](../../pkg/models/shared/script.md)                                                                  | :heavy_check_mark:                                                                                                  | N/A                                                                                                                 |                                                                                                                     |
| `ledger`                                                                                                            | *string*                                                                                                            | :heavy_check_mark:                                                                                                  | Name of the ledger.                                                                                                 | ledger001                                                                                                           |
| `preview`                                                                                                           | **bool*                                                                                                             | :heavy_minus_sign:                                                                                                  | Set the preview mode. Preview mode doesn't add the logs to the database or publish a message to the message broker. | true                                                                                                                |


### Response

**[*operations.RunScriptResponse](../../pkg/models/operations/runscriptresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
