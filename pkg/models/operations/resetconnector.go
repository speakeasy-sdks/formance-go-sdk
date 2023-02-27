package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type ResetConnectorPathParams struct {
	Connector shared.ConnectorEnum `pathParam:"style=simple,explode=false,name=connector"`
}

type ResetConnectorRequest struct {
	PathParams ResetConnectorPathParams
}

type ResetConnectorResponse struct {
	ContentType string
	StatusCode  int
}
