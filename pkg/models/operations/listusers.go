package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type ListUsersResponse struct {
	ContentType       string
	ListUsersResponse *shared.ListUsersResponse
	StatusCode        int
}
