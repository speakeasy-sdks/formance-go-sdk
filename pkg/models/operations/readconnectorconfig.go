package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ReadConnectorConfigRequest struct {
	Connector shared.ConnectorEnum `pathParam:"style=simple,explode=false,name=connector"`
}

type ReadConnectorConfigResponse struct {
	ConnectorConfigResponse *shared.ConnectorConfigResponse
	ContentType             string
	StatusCode              int
	RawResponse             *http.Response
}
