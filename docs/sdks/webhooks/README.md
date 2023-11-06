# Webhooks
(*Webhooks*)

## Overview

Everything related to Webhooks

### Available Operations

* [ActivateConfig](#activateconfig) - Activate one config
* [ChangeConfigSecret](#changeconfigsecret) - Change the signing secret of a config
* [DeactivateConfig](#deactivateconfig) - Deactivate one config
* [DeleteConfig](#deleteconfig) - Delete one config
* [GetManyConfigs](#getmanyconfigs) - Get many configs
* [InsertConfig](#insertconfig) - Insert a new config
* [TestConfig](#testconfig) - Test one config

## ActivateConfig

Activate a webhooks config by ID, to start receiving webhooks to its endpoint.

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


    var id string = "4997257d-dfb6-445b-929c-cbe2ab182818"

    ctx := context.Background()
    res, err := s.Webhooks.ActivateConfig(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Config ID                                             | 4997257d-dfb6-445b-929c-cbe2ab182818                  |


### Response

**[*operations.ActivateConfigResponse](../../models/operations/activateconfigresponse.md), error**


## ChangeConfigSecret

Change the signing secret of the endpoint of a webhooks config.

If not passed or empty, a secret is automatically generated.
The format is a random string of bytes of size 24, base64 encoded. (larger size after encoding)


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


    var id string = "4997257d-dfb6-445b-929c-cbe2ab182818"

    configChangeSecret := &shared.ConfigChangeSecret{
        Secret: formancegosdk.String("V0bivxRWveaoz08afqjU6Ko/jwO0Cb+3"),
    }

    ctx := context.Background()
    res, err := s.Webhooks.ChangeConfigSecret(ctx, id, configChangeSecret)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |                                                                         |
| `id`                                                                    | *string*                                                                | :heavy_check_mark:                                                      | Config ID                                                               | 4997257d-dfb6-445b-929c-cbe2ab182818                                    |
| `configChangeSecret`                                                    | [*shared.ConfigChangeSecret](../../models/shared/configchangesecret.md) | :heavy_minus_sign:                                                      | N/A                                                                     |                                                                         |


### Response

**[*operations.ChangeConfigSecretResponse](../../models/operations/changeconfigsecretresponse.md), error**


## DeactivateConfig

Deactivate a webhooks config by ID, to stop receiving webhooks to its endpoint.

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


    var id string = "4997257d-dfb6-445b-929c-cbe2ab182818"

    ctx := context.Background()
    res, err := s.Webhooks.DeactivateConfig(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Config ID                                             | 4997257d-dfb6-445b-929c-cbe2ab182818                  |


### Response

**[*operations.DeactivateConfigResponse](../../models/operations/deactivateconfigresponse.md), error**


## DeleteConfig

Delete a webhooks config by ID.

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


    var id string = "4997257d-dfb6-445b-929c-cbe2ab182818"

    ctx := context.Background()
    res, err := s.Webhooks.DeleteConfig(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Config ID                                             | 4997257d-dfb6-445b-929c-cbe2ab182818                  |


### Response

**[*operations.DeleteConfigResponse](../../models/operations/deleteconfigresponse.md), error**


## GetManyConfigs

Sorted by updated date descending

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


    var endpoint *string = "https://example.com"

    var id *string = "4997257d-dfb6-445b-929c-cbe2ab182818"

    ctx := context.Background()
    res, err := s.Webhooks.GetManyConfigs(ctx, endpoint, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `endpoint`                                            | **string*                                             | :heavy_minus_sign:                                    | Optional filter by endpoint URL                       | https://example.com                                   |
| `id`                                                  | **string*                                             | :heavy_minus_sign:                                    | Optional filter by Config ID                          | 4997257d-dfb6-445b-929c-cbe2ab182818                  |


### Response

**[*operations.GetManyConfigsResponse](../../models/operations/getmanyconfigsresponse.md), error**


## InsertConfig

Insert a new webhooks config.

The endpoint should be a valid https URL and be unique.

The secret is the endpoint's verification secret.
If not passed or empty, a secret is automatically generated.
The format is a random string of bytes of size 24, base64 encoded. (larger size after encoding)

All eventTypes are converted to lower-case when inserted.


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
    res, err := s.Webhooks.InsertConfig(ctx, shared.ConfigUser{
        Endpoint: "https://example.com",
        EventTypes: []string{
            "TYPE1",
            "TYPE2",
        },
        Secret: formancegosdk.String("V0bivxRWveaoz08afqjU6Ko/jwO0Cb+3"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                              | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |
| `request`                                              | [shared.ConfigUser](../../models/shared/configuser.md) | :heavy_check_mark:                                     | The request object to use for the request.             |


### Response

**[*operations.InsertConfigResponse](../../models/operations/insertconfigresponse.md), error**


## TestConfig

Test a config by sending a webhook to its endpoint.

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


    var id string = "4997257d-dfb6-445b-929c-cbe2ab182818"

    ctx := context.Background()
    res, err := s.Webhooks.TestConfig(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.AttemptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Config ID                                             | 4997257d-dfb6-445b-929c-cbe2ab182818                  |


### Response

**[*operations.TestConfigResponse](../../models/operations/testconfigresponse.md), error**

