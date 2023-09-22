// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/utils"
	"net/http"
)

type GetTransactionsRequest struct {
	// Parameter used in pagination requests.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when the cursor is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// The maximum number of results to return per page
	PageSize *int64 `default:"15" queryParam:"style=form,explode=true,name=pageSize"`
	// A wallet ID to filter on
	WalletID *string `queryParam:"style=form,explode=true,name=wallet_id"`
}

func (g GetTransactionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTransactionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTransactionsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *GetTransactionsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *GetTransactionsRequest) GetWalletID() *string {
	if o == nil {
		return nil
	}
	return o.WalletID
}

type GetTransactionsResponse struct {
	ContentType string
	// OK
	GetTransactionsResponse *shared.GetTransactionsResponse
	StatusCode              int
	RawResponse             *http.Response
	// Error
	WalletsErrorResponse *shared.WalletsErrorResponse
}

func (o *GetTransactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransactionsResponse) GetGetTransactionsResponse() *shared.GetTransactionsResponse {
	if o == nil {
		return nil
	}
	return o.GetTransactionsResponse
}

func (o *GetTransactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTransactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTransactionsResponse) GetWalletsErrorResponse() *shared.WalletsErrorResponse {
	if o == nil {
		return nil
	}
	return o.WalletsErrorResponse
}
