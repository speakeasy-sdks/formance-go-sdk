// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ListScopesResponse struct {
	ContentType string
	// List of scopes
	ListScopesResponse *shared.ListScopesResponse
	StatusCode         int
	RawResponse        *http.Response
}
