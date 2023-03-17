package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ReadUserRequest struct {
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

type ReadUserResponse struct {
	ContentType      string
	ReadUserResponse *shared.ReadUserResponse
	StatusCode       int
	RawResponse      *http.Response
}
