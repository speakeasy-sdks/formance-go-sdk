package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type GetBalancePathParams struct {
	BalanceName string `pathParam:"style=simple,explode=false,name=balanceName"`
	ID          string `pathParam:"style=simple,explode=false,name=id"`
}

type GetBalanceRequest struct {
	PathParams GetBalancePathParams
}

type GetBalanceResponse struct {
	ContentType          string
	GetBalanceResponse   *shared.GetBalanceResponse
	StatusCode           int
	WalletsErrorResponse *shared.WalletsErrorResponse
}
