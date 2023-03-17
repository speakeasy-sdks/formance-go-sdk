package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type AddMetadataOnTransactionRequest struct {
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	Ledger      string                 `pathParam:"style=simple,explode=false,name=ledger"`
	Txid        int64                  `pathParam:"style=simple,explode=false,name=txid"`
}

type AddMetadataOnTransactionResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	RawResponse   *http.Response
}
