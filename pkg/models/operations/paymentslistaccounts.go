package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type PaymentslistAccountsRequest struct {
	Cursor   *string  `queryParam:"style=form,explode=true,name=cursor"`
	PageSize *int64   `queryParam:"style=form,explode=true,name=pageSize"`
	Sort     []string `queryParam:"style=form,explode=true,name=sort"`
}

type PaymentslistAccountsResponse struct {
	AccountsCursor *shared.AccountsCursor
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}
