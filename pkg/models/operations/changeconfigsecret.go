package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type ChangeConfigSecretPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type ChangeConfigSecretRequest struct {
	PathParams ChangeConfigSecretPathParams
	Request    *shared.ConfigChangeSecret `request:"mediaType=application/json"`
}

type ChangeConfigSecretResponse struct {
	ConfigResponse *shared.ConfigResponse
	ContentType    string
	StatusCode     int
}
