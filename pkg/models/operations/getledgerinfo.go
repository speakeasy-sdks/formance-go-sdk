package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetLedgerInfoRequest struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type GetLedgerInfoResponse struct {
	ContentType        string
	ErrorResponse      *shared.ErrorResponse
	LedgerInfoResponse interface{}
	StatusCode         int
	RawResponse        *http.Response
}
