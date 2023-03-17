package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreditWalletRequest struct {
	CreditWalletRequest *shared.CreditWalletRequest `request:"mediaType=application/json"`
	ID                  string                      `pathParam:"style=simple,explode=false,name=id"`
}

type CreditWalletResponse struct {
	ContentType          string
	StatusCode           int
	RawResponse          *http.Response
	WalletsErrorResponse *shared.WalletsErrorResponse
}
