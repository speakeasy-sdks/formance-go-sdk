package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type ConnectorsStripeTransferRequest struct {
	Request shared.StripeTransferRequest `request:"mediaType=application/json"`
}

type ConnectorsStripeTransferResponse struct {
	ContentType            string
	StatusCode             int
	StripeTransferResponse map[string]interface{}
}
