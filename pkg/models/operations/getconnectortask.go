package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetConnectorTaskPathParams struct {
	Connector shared.ConnectorEnum `pathParam:"style=simple,explode=false,name=connector"`
	TaskID    string               `pathParam:"style=simple,explode=false,name=taskId"`
}

type GetConnectorTaskRequest struct {
	PathParams GetConnectorTaskPathParams
}

type GetConnectorTaskResponse struct {
	ContentType  string
	StatusCode   int
	RawResponse  *http.Response
	TaskResponse *shared.TaskResponse
}
