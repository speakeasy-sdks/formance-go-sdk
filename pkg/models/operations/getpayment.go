package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetPaymentRequest struct {
	PaymentID string `pathParam:"style=simple,explode=false,name=paymentId"`
}

type GetPaymentResponse struct {
	ContentType     string
	PaymentResponse *shared.PaymentResponse
	StatusCode      int
	RawResponse     *http.Response
}
