package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type CreateTransactionsPathParams struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type CreateTransactionsRequest struct {
	PathParams CreateTransactionsPathParams
	Request    shared.Transactions `request:"mediaType=application/json"`
}

type CreateTransactionsResponse struct {
	ContentType          string
	ErrorResponse        *shared.ErrorResponse
	StatusCode           int
	TransactionsResponse *shared.TransactionsResponse
}
