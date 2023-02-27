package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type CreateWalletRequest struct {
	Request *shared.CreateWalletRequest `request:"mediaType=application/json"`
}

type CreateWalletResponse struct {
	ContentType          string
	CreateWalletResponse *shared.CreateWalletResponse
	StatusCode           int
	WalletsErrorResponse *shared.WalletsErrorResponse
}
