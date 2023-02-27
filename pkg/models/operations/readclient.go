package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type ReadClientPathParams struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

type ReadClientRequest struct {
	PathParams ReadClientPathParams
}

type ReadClientResponse struct {
	ContentType        string
	ReadClientResponse *shared.ReadClientResponse
	StatusCode         int
}
