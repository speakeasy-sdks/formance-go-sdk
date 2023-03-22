// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ReadScopeRequest struct {
	// Scope ID
	ScopeID string `pathParam:"style=simple,explode=false,name=scopeId"`
}

type ReadScopeResponse struct {
	ContentType string
	// Retrieved scope
	ReadScopeResponse *shared.ReadScopeResponse
	StatusCode        int
	RawResponse       *http.Response
}
