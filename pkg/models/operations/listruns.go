// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ListRunsRequest struct {
	// The flow id
	FlowID string `pathParam:"style=simple,explode=false,name=flowId"`
}

type ListRunsResponse struct {
	ContentType string
	// General error
	Error *shared.Error
	// List of workflow occurrences
	ListRunsResponse *shared.ListRunsResponse
	StatusCode       int
	RawResponse      *http.Response
}
