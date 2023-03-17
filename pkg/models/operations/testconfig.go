package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type TestConfigRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type TestConfigResponse struct {
	AttemptResponse *shared.AttemptResponse
	ContentType     string
	StatusCode      int
	RawResponse     *http.Response
}
