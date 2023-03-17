package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateWorkflowResponse struct {
	ContentType            string
	CreateWorkflowResponse *shared.CreateWorkflowResponse
	Error                  *shared.Error
	StatusCode             int
	RawResponse            *http.Response
}
