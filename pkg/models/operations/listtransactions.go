package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"time"
)

type ListTransactionsPathParams struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type ListTransactionsQueryParams struct {
	Account             *string                `queryParam:"style=form,explode=true,name=account"`
	After               *string                `queryParam:"style=form,explode=true,name=after"`
	Cursor              *string                `queryParam:"style=form,explode=true,name=cursor"`
	Destination         *string                `queryParam:"style=form,explode=true,name=destination"`
	EndTime             *time.Time             `queryParam:"style=form,explode=true,name=endTime"`
	EndTimeDeprecated   *time.Time             `queryParam:"style=form,explode=true,name=end_time"`
	Metadata            map[string]interface{} `queryParam:"style=deepObject,explode=true,name=metadata"`
	PageSize            *int64                 `queryParam:"style=form,explode=true,name=pageSize"`
	PageSizeDeprecated  *int64                 `queryParam:"style=form,explode=true,name=page_size"`
	PaginationToken     *string                `queryParam:"style=form,explode=true,name=pagination_token"`
	Reference           *string                `queryParam:"style=form,explode=true,name=reference"`
	Source              *string                `queryParam:"style=form,explode=true,name=source"`
	StartTime           *time.Time             `queryParam:"style=form,explode=true,name=startTime"`
	StartTimeDeprecated *time.Time             `queryParam:"style=form,explode=true,name=start_time"`
}

type ListTransactionsRequest struct {
	PathParams  ListTransactionsPathParams
	QueryParams ListTransactionsQueryParams
}

type ListTransactionsResponse struct {
	ContentType                string
	ErrorResponse              *shared.ErrorResponse
	StatusCode                 int
	TransactionsCursorResponse *shared.TransactionsCursorResponse
}
