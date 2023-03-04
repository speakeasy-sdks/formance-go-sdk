package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type PaymentsgetServerInfoResponse struct {
	ContentType string
	ServerInfo  *shared.ServerInfo
	StatusCode  int
	RawResponse *http.Response
}
