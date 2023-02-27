package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type CreateScopeRequest struct {
	Request *shared.CreateScopeRequest `request:"mediaType=application/json"`
}

type CreateScopeResponse struct {
	ContentType         string
	CreateScopeResponse *shared.CreateScopeResponse
	StatusCode          int
}
