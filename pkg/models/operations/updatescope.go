package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateScopePathParams struct {
	ScopeID string `pathParam:"style=simple,explode=false,name=scopeId"`
}

type UpdateScopeRequest struct {
	PathParams UpdateScopePathParams
	Request    *shared.UpdateScopeRequest `request:"mediaType=application/json"`
}

type UpdateScopeResponse struct {
	ContentType         string
	StatusCode          int
	RawResponse         *http.Response
	UpdateScopeResponse *shared.UpdateScopeResponse
}
