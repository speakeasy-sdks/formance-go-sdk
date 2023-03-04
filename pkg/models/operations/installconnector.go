package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type InstallConnectorPathParams struct {
	Connector shared.ConnectorEnum `pathParam:"style=simple,explode=false,name=connector"`
}

type InstallConnectorRequest struct {
	PathParams InstallConnectorPathParams
	Request    interface{} `request:"mediaType=application/json"`
}

type InstallConnectorResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
