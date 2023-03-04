package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetHoldPathParams struct {
	HoldID string `pathParam:"style=simple,explode=false,name=holdID"`
}

type GetHoldRequest struct {
	PathParams GetHoldPathParams
}

type GetHoldResponse struct {
	ContentType          string
	GetHoldResponse      *shared.GetHoldResponse
	StatusCode           int
	RawResponse          *http.Response
	WalletsErrorResponse *shared.WalletsErrorResponse
}
