package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type DeactivateConfigRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeactivateConfigResponse struct {
	ConfigResponse *shared.ConfigResponse
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}
