package operations

import (
	"net/http"
)

type DeleteTransientScopeRequest struct {
	ScopeID          string `pathParam:"style=simple,explode=false,name=scopeId"`
	TransientScopeID string `pathParam:"style=simple,explode=false,name=transientScopeId"`
}

type DeleteTransientScopeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
