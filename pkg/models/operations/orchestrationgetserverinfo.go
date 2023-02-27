package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type OrchestrationgetServerInfoResponse struct {
	ContentType string
	Error       *shared.Error
	ServerInfo  *shared.ServerInfo
	StatusCode  int
}
