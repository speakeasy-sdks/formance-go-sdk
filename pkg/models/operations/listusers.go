package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ListUsersResponse struct {
	ContentType       string
	ListUsersResponse *shared.ListUsersResponse
	StatusCode        int
	RawResponse       *http.Response
}
