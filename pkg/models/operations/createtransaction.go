package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type CreateTransactionPathParams struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type CreateTransactionQueryParams struct {
	Preview *bool `queryParam:"style=form,explode=true,name=preview"`
}

type CreateTransactionRequest struct {
	PathParams  CreateTransactionPathParams
	QueryParams CreateTransactionQueryParams
	Request     shared.PostTransaction `request:"mediaType=application/json"`
}

type CreateTransactionResponse struct {
	ContentType          string
	ErrorResponse        *shared.ErrorResponse
	StatusCode           int
	TransactionsResponse *shared.TransactionsResponse
}
