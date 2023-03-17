package operations

import (
	"net/http"
)

type DeleteScopeRequest struct {
	ScopeID string `pathParam:"style=simple,explode=false,name=scopeId"`
}

type DeleteScopeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
