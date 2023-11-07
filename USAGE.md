<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := formancegosdk.New(
		formancegosdk.WithSecurity(""),
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