// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateScopeRequest struct {
	// Scope ID
	ScopeID            string                     `pathParam:"style=simple,explode=false,name=scopeId"`
	UpdateScopeRequest *shared.UpdateScopeRequest `request:"mediaType=application/json"`
}

func (o *UpdateScopeRequest) GetScopeID() string {
	if o == nil {
		return ""
	}
	return o.ScopeID
}

func (o *UpdateScopeRequest) GetUpdateScopeRequest() *shared.UpdateScopeRequest {
	if o == nil {
		return nil
	}
	return o.UpdateScopeRequest
}

type UpdateScopeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updated scope
	UpdateScopeResponse *shared.UpdateScopeResponse
}

func (o *UpdateScopeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateScopeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateScopeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateScopeResponse) GetUpdateScopeResponse() *shared.UpdateScopeResponse {
	if o == nil {
		return nil
	}
	return o.UpdateScopeResponse
}
