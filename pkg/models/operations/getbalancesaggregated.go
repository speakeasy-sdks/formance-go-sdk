package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetBalancesAggregatedRequest struct {
	Address *string `queryParam:"style=form,explode=true,name=address"`
	Ledger  string  `pathParam:"style=simple,explode=false,name=ledger"`
}

type GetBalancesAggregatedResponse struct {
	AggregateBalancesResponse *shared.AggregateBalancesResponse
	ContentType               string
	ErrorResponse             *shared.ErrorResponse
	StatusCode                int
	RawResponse               *http.Response
}
