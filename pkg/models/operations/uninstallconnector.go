package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type UninstallConnectorRequest struct {
	Connector shared.ConnectorEnum `pathParam:"style=simple,explode=false,name=connector"`
}

type UninstallConnectorResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
