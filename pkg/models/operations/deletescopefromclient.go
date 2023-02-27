package operations

type DeleteScopeFromClientPathParams struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
	ScopeID  string `pathParam:"style=simple,explode=false,name=scopeId"`
}

type DeleteScopeFromClientRequest struct {
	PathParams DeleteScopeFromClientPathParams
}

type DeleteScopeFromClientResponse struct {
	ContentType string
	StatusCode  int
}
