package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ReadClientRequest struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

type ReadClientResponse struct {
	ContentType        string
	ReadClientResponse *shared.ReadClientResponse
	StatusCode         int
	RawResponse        *http.Response
}
