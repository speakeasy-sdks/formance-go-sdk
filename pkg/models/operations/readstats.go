package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ReadStatsRequest struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type ReadStatsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	StatsResponse *shared.StatsResponse
	StatusCode    int
	RawResponse   *http.Response
}
