package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ListClientsResponse struct {
	ContentType         string
	ListClientsResponse *shared.ListClientsResponse
	StatusCode          int
	RawResponse         *http.Response
}
