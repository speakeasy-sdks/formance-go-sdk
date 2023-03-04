package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateClientRequest struct {
	Request *shared.CreateClientRequest `request:"mediaType=application/json"`
}

type CreateClientResponse struct {
	ContentType          string
	CreateClientResponse *shared.CreateClientResponse
	StatusCode           int
	RawResponse          *http.Response
}
