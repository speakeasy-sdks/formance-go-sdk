package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type RunWorkflowRequest struct {
	RequestBody map[string]string `request:"mediaType=application/json"`
	FlowID      string            `pathParam:"style=simple,explode=false,name=flowId"`
	Wait        *bool             `queryParam:"style=form,explode=true,name=wait"`
}

type RunWorkflowResponse struct {
	ContentType         string
	Error               *shared.Error
	RunWorkflowResponse *shared.RunWorkflowResponse
	StatusCode          int
	RawResponse         *http.Response
}
