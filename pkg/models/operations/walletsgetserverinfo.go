package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type WalletsgetServerInfoResponse struct {
	ContentType          string
	ServerInfo           *shared.ServerInfo
	StatusCode           int
	WalletsErrorResponse *shared.WalletsErrorResponse
}
