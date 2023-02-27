package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type GetBalancesAggregatedPathParams struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type GetBalancesAggregatedQueryParams struct {
	Address *string `queryParam:"style=form,explode=true,name=address"`
}

type GetBalancesAggregatedRequest struct {
	PathParams  GetBalancesAggregatedPathParams
	QueryParams GetBalancesAggregatedQueryParams
}

type GetBalancesAggregatedResponse struct {
	AggregateBalancesResponse *shared.AggregateBalancesResponse
	ContentType               string
	ErrorResponse             *shared.ErrorResponse
	StatusCode                int
}
