package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateClientResponse struct {
	ContentType          string
	CreateClientResponse *shared.CreateClientResponse
	StatusCode           int
	RawResponse          *http.Response
}
