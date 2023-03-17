package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
	"time"
)

type ListLogsRequest struct {
	After               *string    `queryParam:"style=form,explode=true,name=after"`
	Cursor              *string    `queryParam:"style=form,explode=true,name=cursor"`
	EndTime             *time.Time `queryParam:"style=form,explode=true,name=endTime"`
	EndTimeDeprecated   *time.Time `queryParam:"style=form,explode=true,name=end_time"`
	Ledger              string     `pathParam:"style=simple,explode=false,name=ledger"`
	PageSize            *int64     `queryParam:"style=form,explode=true,name=pageSize"`
	PageSizeDeprecated  *int64     `queryParam:"style=form,explode=true,name=page_size"`
	PaginationToken     *string    `queryParam:"style=form,explode=true,name=pagination_token"`
	StartTime           *time.Time `queryParam:"style=form,explode=true,name=startTime"`
	StartTimeDeprecated *time.Time `queryParam:"style=form,explode=true,name=start_time"`
}

type ListLogsResponse struct {
	ContentType        string
	ErrorResponse      *shared.ErrorResponse
	LogsCursorResponse *shared.LogsCursorResponse
	StatusCode         int
	RawResponse        *http.Response
}
