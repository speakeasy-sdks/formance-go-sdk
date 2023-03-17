package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ConfirmHoldRequest struct {
	ConfirmHoldRequest *shared.ConfirmHoldRequest `request:"mediaType=application/json"`
	HoldID             string                     `pathParam:"style=simple,explode=false,name=hold_id"`
}

type ConfirmHoldResponse struct {
	ContentType          string
	StatusCode           int
	RawResponse          *http.Response
	WalletsErrorResponse *shared.WalletsErrorResponse
}
