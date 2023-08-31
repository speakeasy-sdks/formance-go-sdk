// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
	"time"
)

type ListLogsRequest struct {
	// Pagination cursor, will return the logs after a given ID. (in descending order).
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Filter transactions that occurred before this timestamp.
	// The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
	//
	EndTime *time.Time `queryParam:"style=form,explode=true,name=endTime"`
	// Filter transactions that occurred before this timestamp.
	// The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
	// Deprecated, please use `endTime` instead.
	//
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	EndTimeDeprecated *time.Time `queryParam:"style=form,explode=true,name=end_time"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// The maximum number of results to return per page.
	//
	PageSize *int64 `queryParam:"style=form,explode=true,name=pageSize"`
	// The maximum number of results to return per page.
	// Deprecated, please use `pageSize` instead.
	//
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	PageSizeDeprecated *int64 `queryParam:"style=form,explode=true,name=page_size"`
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	// Deprecated, please use `cursor` instead.
	//
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	PaginationToken *string `queryParam:"style=form,explode=true,name=pagination_token"`
	// Filter transactions that occurred after this timestamp.
	// The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
	//
	StartTime *time.Time `queryParam:"style=form,explode=true,name=startTime"`
	// Filter transactions that occurred after this timestamp.
	// The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
	// Deprecated, please use `startTime` instead.
	//
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	StartTimeDeprecated *time.Time `queryParam:"style=form,explode=true,name=start_time"`
}

func (o *ListLogsRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListLogsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListLogsRequest) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *ListLogsRequest) GetEndTimeDeprecated() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTimeDeprecated
}

func (o *ListLogsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *ListLogsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListLogsRequest) GetPageSizeDeprecated() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSizeDeprecated
}

func (o *ListLogsRequest) GetPaginationToken() *string {
	if o == nil {
		return nil
	}
	return o.PaginationToken
}

func (o *ListLogsRequest) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *ListLogsRequest) GetStartTimeDeprecated() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTimeDeprecated
}

type ListLogsResponse struct {
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	// OK
	LogsCursorResponse *shared.LogsCursorResponse
	StatusCode         int
	RawResponse        *http.Response
}

func (o *ListLogsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListLogsResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *ListLogsResponse) GetLogsCursorResponse() *shared.LogsCursorResponse {
	if o == nil {
		return nil
	}
	return o.LogsCursorResponse
}

func (o *ListLogsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListLogsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
