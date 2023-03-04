package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ListConfigsAvailableConnectorsResponse struct {
	ConnectorsConfigsResponse *shared.ConnectorsConfigsResponse
	ContentType               string
	StatusCode                int
	RawResponse               *http.Response
}
