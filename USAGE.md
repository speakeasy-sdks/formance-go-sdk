<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/formance-go-sdk"
    "github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
)

func main() {
    s := formance.New(
        WithSecurity(        shared.Security{
            Authorization: shared.SchemeAuthorization{
                Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.GetServerInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerInfo != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->