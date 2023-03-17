package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateSecretRequest struct {
	CreateSecretRequest *shared.CreateSecretRequest `request:"mediaType=application/json"`
	ClientID            string                      `pathParam:"style=simple,explode=false,name=clientId"`
}

type CreateSecretResponse struct {
	ContentType          string
	CreateSecretResponse *shared.CreateSecretResponse
	StatusCode           int
	RawResponse          *http.Response
}
