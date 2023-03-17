package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateTransactionsRequest struct {
	Transactions shared.Transactions `request:"mediaType=application/json"`
	Ledger       string              `pathParam:"style=simple,explode=false,name=ledger"`
}

type CreateTransactionsResponse struct {
	ContentType          string
	ErrorResponse        *shared.ErrorResponse
	StatusCode           int
	RawResponse          *http.Response
	TransactionsResponse *shared.TransactionsResponse
}
