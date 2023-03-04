package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetPaymentPathParams struct {
	PaymentID string `pathParam:"style=simple,explode=false,name=paymentId"`
}

type GetPaymentRequest struct {
	PathParams GetPaymentPathParams
}

type GetPaymentResponse struct {
	ContentType     string
	PaymentResponse *shared.PaymentResponse
	StatusCode      int
	RawResponse     *http.Response
}
