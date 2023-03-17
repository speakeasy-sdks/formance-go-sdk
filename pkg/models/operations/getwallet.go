package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetWalletRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetWalletResponse struct {
	ContentType          string
	GetWalletResponse    *shared.GetWalletResponse
	StatusCode           int
	RawResponse          *http.Response
	WalletsErrorResponse *shared.WalletsErrorResponse
}
