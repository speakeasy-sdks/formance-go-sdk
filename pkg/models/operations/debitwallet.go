package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type DebitWalletPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DebitWalletRequest struct {
	PathParams DebitWalletPathParams
	Request    *shared.DebitWalletRequest `request:"mediaType=application/json"`
}

type DebitWalletResponse struct {
	ContentType          string
	DebitWalletResponse  *shared.DebitWalletResponse
	StatusCode           int
	RawResponse          *http.Response
	WalletsErrorResponse *shared.WalletsErrorResponse
}
