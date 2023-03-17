package operations

import (
	"net/http"
)

type DeleteScopeFromClientRequest struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
	ScopeID  string `pathParam:"style=simple,explode=false,name=scopeId"`
}

type DeleteScopeFromClientResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
