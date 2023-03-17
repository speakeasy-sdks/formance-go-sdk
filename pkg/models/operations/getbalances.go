package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetBalancesRequest struct {
	Address         *string `queryParam:"style=form,explode=true,name=address"`
	After           *string `queryParam:"style=form,explode=true,name=after"`
	Cursor          *string `queryParam:"style=form,explode=true,name=cursor"`
	Ledger          string  `pathParam:"style=simple,explode=false,name=ledger"`
	PaginationToken *string `queryParam:"style=form,explode=true,name=pagination_token"`
}

type GetBalancesResponse struct {
	BalancesCursorResponse *shared.BalancesCursorResponse
	ContentType            string
	ErrorResponse          *shared.ErrorResponse
	StatusCode             int
	RawResponse            *http.Response
}
