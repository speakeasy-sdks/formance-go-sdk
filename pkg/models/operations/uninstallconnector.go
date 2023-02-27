package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type UninstallConnectorPathParams struct {
	Connector shared.ConnectorEnum `pathParam:"style=simple,explode=false,name=connector"`
}

type UninstallConnectorRequest struct {
	PathParams UninstallConnectorPathParams
}

type UninstallConnectorResponse struct {
	ContentType string
	StatusCode  int
}
