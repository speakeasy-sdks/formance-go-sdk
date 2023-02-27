package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type UpdateClientPathParams struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

type UpdateClientRequest struct {
	PathParams UpdateClientPathParams
	Request    *shared.UpdateClientRequest `request:"mediaType=application/json"`
}

type UpdateClientResponse struct {
	ContentType          string
	StatusCode           int
	UpdateClientResponse *shared.UpdateClientResponse
}
