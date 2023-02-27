package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type ListConfigsAvailableConnectorsResponse struct {
	ConnectorsConfigsResponse *shared.ConnectorsConfigsResponse
	ContentType               string
	StatusCode                int
}
