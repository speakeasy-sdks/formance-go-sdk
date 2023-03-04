package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateBalancePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type CreateBalanceRequest struct {
	PathParams CreateBalancePathParams
	Request    *shared.CreateBalanceRequest `request:"mediaType=application/json"`
}

type CreateBalanceResponse struct {
	ContentType           string
	CreateBalanceResponse *shared.CreateBalanceResponse
	StatusCode            int
	RawResponse           *http.Response
	WalletsErrorResponse  *shared.WalletsErrorResponse
}
