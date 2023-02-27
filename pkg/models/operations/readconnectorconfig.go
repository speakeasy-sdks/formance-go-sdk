package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type ReadConnectorConfigPathParams struct {
	Connector shared.ConnectorEnum `pathParam:"style=simple,explode=false,name=connector"`
}

type ReadConnectorConfigRequest struct {
	PathParams ReadConnectorConfigPathParams
}

type ReadConnectorConfigResponse struct {
	ConnectorConfigResponse *shared.ConnectorConfigResponse
	ContentType             string
	StatusCode              int
}
