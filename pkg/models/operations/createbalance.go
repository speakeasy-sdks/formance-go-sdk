package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateBalanceRequest struct {
	CreateBalanceRequest *shared.CreateBalanceRequest `request:"mediaType=application/json"`
	ID                   string                       `pathParam:"style=simple,explode=false,name=id"`
}

type CreateBalanceResponse struct {
	ContentType           string
	CreateBalanceResponse *shared.CreateBalanceResponse
	StatusCode            int
	RawResponse           *http.Response
	WalletsErrorResponse  *shared.WalletsErrorResponse
}
