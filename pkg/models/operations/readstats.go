// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ReadStatsRequest struct {
	// name of the ledger
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type ReadStatsResponse struct {
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	// OK
	StatsResponse *shared.StatsResponse
	StatusCode    int
	RawResponse   *http.Response
}
