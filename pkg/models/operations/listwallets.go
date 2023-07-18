// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

// ListWalletsMetadata - Filter wallets by metadata key value pairs. Nested objects can be used as seen in the example below.
type ListWalletsMetadata struct {
}

type ListWalletsRequest struct {
	// Parameter used in pagination requests.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when the pagination token is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Filter wallets by metadata key value pairs. Nested objects can be used as seen in the example below.
	Metadata *ListWalletsMetadata `queryParam:"style=deepObject,explode=true,name=metadata"`
	// Filter on wallet name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// The maximum number of results to return per page
	PageSize *int64 `queryParam:"style=form,explode=true,name=pageSize"`
}

func (o *ListWalletsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListWalletsRequest) GetMetadata() *ListWalletsMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *ListWalletsRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListWalletsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type ListWalletsResponse struct {
	ContentType string
	// OK
	ListWalletsResponse *shared.ListWalletsResponse
	StatusCode          int
	RawResponse         *http.Response
}

func (o *ListWalletsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListWalletsResponse) GetListWalletsResponse() *shared.ListWalletsResponse {
	if o == nil {
		return nil
	}
	return o.ListWalletsResponse
}

func (o *ListWalletsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListWalletsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
