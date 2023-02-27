package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type PaymentslistAccountsQueryParams struct {
	Cursor   *string  `queryParam:"style=form,explode=true,name=cursor"`
	PageSize *int64   `queryParam:"style=form,explode=true,name=pageSize"`
	Sort     []string `queryParam:"style=form,explode=true,name=sort"`
}

type PaymentslistAccountsRequest struct {
	QueryParams PaymentslistAccountsQueryParams
}

type PaymentslistAccountsResponse struct {
	AccountsCursor *shared.AccountsCursor
	ContentType    string
	StatusCode     int
}
