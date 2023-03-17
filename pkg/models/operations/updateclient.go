package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateClientRequest struct {
	UpdateClientRequest *shared.UpdateClientRequest `request:"mediaType=application/json"`
	ClientID            string                      `pathParam:"style=simple,explode=false,name=clientId"`
}

type UpdateClientResponse struct {
	ContentType          string
	StatusCode           int
	RawResponse          *http.Response
	UpdateClientResponse *shared.UpdateClientResponse
}
