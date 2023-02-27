package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type ReadStatsPathParams struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type ReadStatsRequest struct {
	PathParams ReadStatsPathParams
}

type ReadStatsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	StatsResponse *shared.StatsResponse
	StatusCode    int
}
