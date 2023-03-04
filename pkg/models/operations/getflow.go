package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetFlowPathParams struct {
	FlowID string `pathParam:"style=simple,explode=false,name=flowId"`
}

type GetFlowRequest struct {
	PathParams GetFlowPathParams
}

type GetFlowResponse struct {
	ContentType         string
	Error               *shared.Error
	GetWorkflowResponse *shared.GetWorkflowResponse
	StatusCode          int
	RawResponse         *http.Response
}
