package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateWalletRequestBody struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateWalletRequest struct {
	RequestBody *UpdateWalletRequestBody `request:"mediaType=application/json"`
	ID          string                   `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateWalletResponse struct {
	ContentType          string
	StatusCode           int
	RawResponse          *http.Response
	WalletsErrorResponse *shared.WalletsErrorResponse
}
