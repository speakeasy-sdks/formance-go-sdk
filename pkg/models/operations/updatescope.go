package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateScopeRequest struct {
	UpdateScopeRequest *shared.UpdateScopeRequest `request:"mediaType=application/json"`
	ScopeID            string                     `pathParam:"style=simple,explode=false,name=scopeId"`
}

type UpdateScopeResponse struct {
	ContentType         string
	StatusCode          int
	RawResponse         *http.Response
	UpdateScopeResponse *shared.UpdateScopeResponse
}
