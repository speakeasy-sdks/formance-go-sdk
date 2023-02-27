package operations

type DeleteScopePathParams struct {
	ScopeID string `pathParam:"style=simple,explode=false,name=scopeId"`
}

type DeleteScopeRequest struct {
	PathParams DeleteScopePathParams
}

type DeleteScopeResponse struct {
	ContentType string
	StatusCode  int
}
