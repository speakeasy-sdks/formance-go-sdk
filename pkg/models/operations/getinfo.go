package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetInfoResponse struct {
	ConfigInfoResponse *shared.ConfigInfoResponse
	ContentType        string
	ErrorResponse      *shared.ErrorResponse
	StatusCode         int
	RawResponse        *http.Response
}
