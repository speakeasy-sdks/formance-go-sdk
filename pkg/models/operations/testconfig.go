package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type TestConfigPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type TestConfigRequest struct {
	PathParams TestConfigPathParams
}

type TestConfigResponse struct {
	AttemptResponse *shared.AttemptResponse
	ContentType     string
	StatusCode      int
}
