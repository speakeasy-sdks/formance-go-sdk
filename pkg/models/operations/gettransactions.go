package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetTransactionsQueryParams struct {
	Cursor   *string `queryParam:"style=form,explode=true,name=cursor"`
	PageSize *int64  `queryParam:"style=form,explode=true,name=pageSize"`
	WalletID *string `queryParam:"style=form,explode=true,name=wallet_id"`
}

type GetTransactionsRequest struct {
	QueryParams GetTransactionsQueryParams
}

type GetTransactionsResponse struct {
	ContentType             string
	GetTransactionsResponse *shared.GetTransactionsResponse
	StatusCode              int
	RawResponse             *http.Response
	WalletsErrorResponse    *shared.WalletsErrorResponse
}
