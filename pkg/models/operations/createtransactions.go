// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateTransactionsRequest struct {
	Transactions shared.Transactions `request:"mediaType=application/json"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

func (o *CreateTransactionsRequest) GetTransactions() shared.Transactions {
	if o == nil {
		return shared.Transactions{}
	}
	return o.Transactions
}

func (o *CreateTransactionsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

type CreateTransactionsResponse struct {
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	// OK
	TransactionsResponse *shared.TransactionsResponse
}

func (o *CreateTransactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTransactionsResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *CreateTransactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTransactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTransactionsResponse) GetTransactionsResponse() *shared.TransactionsResponse {
	if o == nil {
		return nil
	}
	return o.TransactionsResponse
}
