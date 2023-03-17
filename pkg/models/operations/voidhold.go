package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type VoidHoldRequest struct {
	HoldID string `pathParam:"style=simple,explode=false,name=hold_id"`
}

type VoidHoldResponse struct {
	ContentType          string
	StatusCode           int
	RawResponse          *http.Response
	WalletsErrorResponse *shared.WalletsErrorResponse
}
