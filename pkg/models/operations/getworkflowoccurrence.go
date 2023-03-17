package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetWorkflowOccurrenceRequest struct {
	FlowID string `pathParam:"style=simple,explode=false,name=flowId"`
	RunID  string `pathParam:"style=simple,explode=false,name=runId"`
}

type GetWorkflowOccurrenceResponse struct {
	ContentType                   string
	Error                         *shared.Error
	GetWorkflowOccurrenceResponse *shared.GetWorkflowOccurrenceResponse
	StatusCode                    int
	RawResponse                   *http.Response
}
