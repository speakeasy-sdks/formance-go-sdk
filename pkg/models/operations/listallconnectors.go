package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ListAllConnectorsResponse struct {
	ConnectorsResponse *shared.ConnectorsResponse
	ContentType        string
	StatusCode         int
	RawResponse        *http.Response
}
