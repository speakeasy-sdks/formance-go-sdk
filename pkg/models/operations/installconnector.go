package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type InstallConnectorRequest struct {
	RequestBody interface{}          `request:"mediaType=application/json"`
	Connector   shared.ConnectorEnum `pathParam:"style=simple,explode=false,name=connector"`
}

type InstallConnectorResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
