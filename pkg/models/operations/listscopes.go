// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ListScopesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of scopes
	ListScopesResponse *shared.ListScopesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListScopesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListScopesResponse) GetListScopesResponse() *shared.ListScopesResponse {
	if o == nil {
		return nil
	}
	return o.ListScopesResponse
}

func (o *ListScopesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListScopesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
