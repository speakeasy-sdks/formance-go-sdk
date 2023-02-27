package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"time"
)

type ListLogsPathParams struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type ListLogsQueryParams struct {
	After               *string    `queryParam:"style=form,explode=true,name=after"`
	Cursor              *string    `queryParam:"style=form,explode=true,name=cursor"`
	EndTime             *time.Time `queryParam:"style=form,explode=true,name=endTime"`
	EndTimeDeprecated   *time.Time `queryParam:"style=form,explode=true,name=end_time"`
	PageSize            *int64     `queryParam:"style=form,explode=true,name=pageSize"`
	PageSizeDeprecated  *int64     `queryParam:"style=form,explode=true,name=page_size"`
	PaginationToken     *string    `queryParam:"style=form,explode=true,name=pagination_token"`
	StartTime           *time.Time `queryParam:"style=form,explode=true,name=startTime"`
	StartTimeDeprecated *time.Time `queryParam:"style=form,explode=true,name=start_time"`
}

type ListLogsRequest struct {
	PathParams  ListLogsPathParams
	QueryParams ListLogsQueryParams
}

type ListLogsResponse struct {
	ContentType        string
	ErrorResponse      *shared.ErrorResponse
	LogsCursorResponse *shared.LogsCursorResponse
	StatusCode         int
}
