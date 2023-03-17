package operations

import (
	"net/http"
)

type AddScopeToClientRequest struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
	ScopeID  string `pathParam:"style=simple,explode=false,name=scopeId"`
}

type AddScopeToClientResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
