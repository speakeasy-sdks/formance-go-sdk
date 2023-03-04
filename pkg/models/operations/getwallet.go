package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetWalletPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetWalletRequest struct {
	PathParams GetWalletPathParams
}

type GetWalletResponse struct {
	ContentType          string
	GetWalletResponse    *shared.GetWalletResponse
	StatusCode           int
	RawResponse          *http.Response
	WalletsErrorResponse *shared.WalletsErrorResponse
}
