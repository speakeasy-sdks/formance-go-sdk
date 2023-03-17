package operations

import (
	"net/http"
)

type AddTransientScopeRequest struct {
	ScopeID          string `pathParam:"style=simple,explode=false,name=scopeId"`
	TransientScopeID string `pathParam:"style=simple,explode=false,name=transientScopeId"`
}

type AddTransientScopeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
