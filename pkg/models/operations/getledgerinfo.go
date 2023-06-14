// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetLedgerInfoRequest struct {
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type GetLedgerInfoResponse struct {
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	// OK
	LedgerInfoResponse *shared.LedgerInfoResponse
	StatusCode         int
	RawResponse        *http.Response
}
