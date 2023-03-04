package operations

import (
	"net/http"
)

type AddScopeToClientPathParams struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
	ScopeID  string `pathParam:"style=simple,explode=false,name=scopeId"`
}

type AddScopeToClientRequest struct {
	PathParams AddScopeToClientPathParams
}

type AddScopeToClientResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
