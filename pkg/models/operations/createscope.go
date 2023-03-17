package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateScopeResponse struct {
	ContentType         string
	CreateScopeResponse *shared.CreateScopeResponse
	StatusCode          int
	RawResponse         *http.Response
}
