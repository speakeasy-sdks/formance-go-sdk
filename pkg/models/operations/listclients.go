package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type ListClientsResponse struct {
	ContentType         string
	ListClientsResponse *shared.ListClientsResponse
	StatusCode          int
}
