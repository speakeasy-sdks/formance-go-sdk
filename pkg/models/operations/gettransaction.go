package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type GetTransactionPathParams struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	Txid   int64  `pathParam:"style=simple,explode=false,name=txid"`
}

type GetTransactionRequest struct {
	PathParams GetTransactionPathParams
}

type GetTransactionResponse struct {
	ContentType         string
	ErrorResponse       *shared.ErrorResponse
	StatusCode          int
	TransactionResponse *shared.TransactionResponse
}
