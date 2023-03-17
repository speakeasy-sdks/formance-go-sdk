package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ChangeConfigSecretRequest struct {
	ConfigChangeSecret *shared.ConfigChangeSecret `request:"mediaType=application/json"`
	ID                 string                     `pathParam:"style=simple,explode=false,name=id"`
}

type ChangeConfigSecretResponse struct {
	ConfigResponse *shared.ConfigResponse
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}
