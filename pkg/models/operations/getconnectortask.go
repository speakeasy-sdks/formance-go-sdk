package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetConnectorTaskRequest struct {
	Connector shared.ConnectorEnum `pathParam:"style=simple,explode=false,name=connector"`
	TaskID    string               `pathParam:"style=simple,explode=false,name=taskId"`
}

type GetConnectorTaskResponse struct {
	ContentType  string
	StatusCode   int
	RawResponse  *http.Response
	TaskResponse *shared.TaskResponse
}
