package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateWorkflowRequest struct {
	Request *shared.CreateWorkflowRequest `request:"mediaType=application/json"`
}

type CreateWorkflowResponse struct {
	ContentType            string
	CreateWorkflowResponse *shared.CreateWorkflowResponse
	Error                  *shared.Error
	StatusCode             int
	RawResponse            *http.Response
}
