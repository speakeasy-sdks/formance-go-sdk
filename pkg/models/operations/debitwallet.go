package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type DebitWalletRequest struct {
	DebitWalletRequest *shared.DebitWalletRequest `request:"mediaType=application/json"`
	ID                 string                     `pathParam:"style=simple,explode=false,name=id"`
}

type DebitWalletResponse struct {
	ContentType          string
	DebitWalletResponse  *shared.DebitWalletResponse
	StatusCode           int
	RawResponse          *http.Response
	WalletsErrorResponse *shared.WalletsErrorResponse
}
