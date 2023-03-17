package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ReadScopeRequest struct {
	ScopeID string `pathParam:"style=simple,explode=false,name=scopeId"`
}

type ReadScopeResponse struct {
	ContentType       string
	ReadScopeResponse *shared.ReadScopeResponse
	StatusCode        int
	RawResponse       *http.Response
}
