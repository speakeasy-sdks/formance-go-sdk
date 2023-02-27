package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type UpdateWalletPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateWalletRequestBody struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateWalletRequest struct {
	PathParams UpdateWalletPathParams
	Request    *UpdateWalletRequestBody `request:"mediaType=application/json"`
}

type UpdateWalletResponse struct {
	ContentType          string
	StatusCode           int
	WalletsErrorResponse *shared.WalletsErrorResponse
}
