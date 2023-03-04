package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type RunWorkflowPathParams struct {
	FlowID string `pathParam:"style=simple,explode=false,name=flowId"`
}

type RunWorkflowQueryParams struct {
	Wait *bool `queryParam:"style=form,explode=true,name=wait"`
}

type RunWorkflowRequest struct {
	PathParams  RunWorkflowPathParams
	QueryParams RunWorkflowQueryParams
	Request     map[string]string `request:"mediaType=application/json"`
}

type RunWorkflowResponse struct {
	ContentType         string
	Error               *shared.Error
	RunWorkflowResponse *shared.RunWorkflowResponse
	StatusCode          int
	RawResponse         *http.Response
}
