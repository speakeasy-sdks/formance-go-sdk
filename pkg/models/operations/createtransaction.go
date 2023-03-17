package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateTransactionRequest struct {
	PostTransaction shared.PostTransaction `request:"mediaType=application/json"`
	Ledger          string                 `pathParam:"style=simple,explode=false,name=ledger"`
	Preview         *bool                  `queryParam:"style=form,explode=true,name=preview"`
}

type CreateTransactionResponse struct {
	ContentType          string
	ErrorResponse        *shared.ErrorResponse
	StatusCode           int
	RawResponse          *http.Response
	TransactionsResponse *shared.TransactionsResponse
}
