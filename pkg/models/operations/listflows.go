package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ListFlowsResponse struct {
	ContentType           string
	Error                 *shared.Error
	ListWorkflowsResponse *shared.ListWorkflowsResponse
	StatusCode            int
	RawResponse           *http.Response
}
