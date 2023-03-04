package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetBalancesPathParams struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type GetBalancesQueryParams struct {
	Address         *string `queryParam:"style=form,explode=true,name=address"`
	After           *string `queryParam:"style=form,explode=true,name=after"`
	Cursor          *string `queryParam:"style=form,explode=true,name=cursor"`
	PaginationToken *string `queryParam:"style=form,explode=true,name=pagination_token"`
}

type GetBalancesRequest struct {
	PathParams  GetBalancesPathParams
	QueryParams GetBalancesQueryParams
}

type GetBalancesResponse struct {
	BalancesCursorResponse *shared.BalancesCursorResponse
	ContentType            string
	ErrorResponse          *shared.ErrorResponse
	StatusCode             int
	RawResponse            *http.Response
}
