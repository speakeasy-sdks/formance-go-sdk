package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ConfirmHoldPathParams struct {
	HoldID string `pathParam:"style=simple,explode=false,name=hold_id"`
}

type ConfirmHoldRequest struct {
	PathParams ConfirmHoldPathParams
	Request    *shared.ConfirmHoldRequest `request:"mediaType=application/json"`
}

type ConfirmHoldResponse struct {
	ContentType          string
	StatusCode           int
	RawResponse          *http.Response
	WalletsErrorResponse *shared.WalletsErrorResponse
}
