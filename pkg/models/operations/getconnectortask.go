// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetConnectorTaskRequest struct {
	// The name of the connector.
	Connector shared.ConnectorEnum `pathParam:"style=simple,explode=false,name=connector"`
	// The task ID.
	TaskID string `pathParam:"style=simple,explode=false,name=taskId"`
}

type GetConnectorTaskResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	TaskResponse *shared.TaskResponse
}
