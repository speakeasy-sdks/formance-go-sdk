package operations

import (
	"net/http"
)

type AddTransientScopePathParams struct {
	ScopeID          string `pathParam:"style=simple,explode=false,name=scopeId"`
	TransientScopeID string `pathParam:"style=simple,explode=false,name=transientScopeId"`
}

type AddTransientScopeRequest struct {
	PathParams AddTransientScopePathParams
}

type AddTransientScopeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
