package operations

import (
	"net/http"
)

type ConnectorsStripeTransferResponse struct {
	ContentType            string
	StatusCode             int
	RawResponse            *http.Response
	StripeTransferResponse map[string]interface{}
}
