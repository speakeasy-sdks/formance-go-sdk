# Orchestration

## Overview

Everything related to Orchestration

### Available Operations

* [CreateWorkflow](#createworkflow) - Create workflow
* [GetFlow](#getflow) - Get a flow by id
* [GetWorkflowOccurrence](#getworkflowoccurrence) - Get a workflow occurrence by id
* [ListFlows](#listflows) - List registered flows
* [ListRuns](#listruns) - List occurrences of a workflow
* [OrchestrationgetServerInfo](#orchestrationgetserverinfo) - Get server info
* [RunWorkflow](#runworkflow) - Run workflow

## CreateWorkflow

Create a workflow

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
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.CreateWorkflow(ctx, shared.CreateWorkflowRequest{
        Stages: []map[string]interface{}{
            map[string]interface{}{
                "saepe": "quidem",
                "architecto": "ipsa",
            },
            map[string]interface{}{
                "est": "mollitia",
                "laborum": "dolores",
                "dolorem": "corporis",
                "explicabo": "nobis",
            },
            map[string]interface{}{
                "omnis": "nemo",
                "minima": "excepturi",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateWorkflowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.CreateWorkflowRequest](../../models/shared/createworkflowrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.CreateWorkflowResponse](../../models/operations/createworkflowresponse.md), error**


## GetFlow

Get a flow by id

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
    res, err := s.Orchestration.GetFlow(ctx, operations.GetFlowRequest{
        FlowID: "accusantium",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetWorkflowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetFlowRequest](../../models/operations/getflowrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.GetFlowResponse](../../models/operations/getflowresponse.md), error**


## GetWorkflowOccurrence

Get a workflow occurrence by id

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
    res, err := s.Orchestration.GetWorkflowOccurrence(ctx, operations.GetWorkflowOccurrenceRequest{
        FlowID: "iure",
        RunID: "culpa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetWorkflowOccurrenceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetWorkflowOccurrenceRequest](../../models/operations/getworkflowoccurrencerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetWorkflowOccurrenceResponse](../../models/operations/getworkflowoccurrenceresponse.md), error**


## ListFlows

List registered flows

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
    res, err := s.Orchestration.ListFlows(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListWorkflowsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListFlowsResponse](../../models/operations/listflowsresponse.md), error**


## ListRuns

List occurrences of a workflow

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
    res, err := s.Orchestration.ListRuns(ctx, operations.ListRunsRequest{
        FlowID: "doloribus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListRunsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.ListRunsRequest](../../models/operations/listrunsrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.ListRunsResponse](../../models/operations/listrunsresponse.md), error**


## OrchestrationgetServerInfo

Get server info

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
    res, err := s.Orchestration.OrchestrationgetServerInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.OrchestrationgetServerInfoResponse](../../models/operations/orchestrationgetserverinforesponse.md), error**


## RunWorkflow

Run workflow

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
    res, err := s.Orchestration.RunWorkflow(ctx, operations.RunWorkflowRequest{
        RequestBody: map[string]string{
            "architecto": "mollitia",
            "dolorem": "culpa",
            "consequuntur": "repellat",
            "mollitia": "occaecati",
        },
        FlowID: "numquam",
        Wait: formance.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RunWorkflowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.RunWorkflowRequest](../../models/operations/runworkflowrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.RunWorkflowResponse](../../models/operations/runworkflowresponse.md), error**
