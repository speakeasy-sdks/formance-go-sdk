package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type CreateSecretPathParams struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

type CreateSecretRequest struct {
	PathParams CreateSecretPathParams
	Request    *shared.CreateSecretRequest `request:"mediaType=application/json"`
}

type CreateSecretResponse struct {
	ContentType          string
	CreateSecretResponse *shared.CreateSecretResponse
	StatusCode           int
}
