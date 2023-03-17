package operations

import (
	"net/http"
)

type DeleteSecretRequest struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
	SecretID string `pathParam:"style=simple,explode=false,name=secretId"`
}

type DeleteSecretResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
