package operations

type DeleteTransientScopePathParams struct {
	ScopeID          string `pathParam:"style=simple,explode=false,name=scopeId"`
	TransientScopeID string `pathParam:"style=simple,explode=false,name=transientScopeId"`
}

type DeleteTransientScopeRequest struct {
	PathParams DeleteTransientScopePathParams
}

type DeleteTransientScopeResponse struct {
	ContentType string
	StatusCode  int
}
