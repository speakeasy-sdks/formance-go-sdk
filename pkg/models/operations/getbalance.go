package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetBalanceRequest struct {
	BalanceName string `pathParam:"style=simple,explode=false,name=balanceName"`
	ID          string `pathParam:"style=simple,explode=false,name=id"`
}

type GetBalanceResponse struct {
	ContentType          string
	GetBalanceResponse   *shared.GetBalanceResponse
	StatusCode           int
	RawResponse          *http.Response
	WalletsErrorResponse *shared.WalletsErrorResponse
}
