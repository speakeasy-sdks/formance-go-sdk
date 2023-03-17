package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type AddMetadataToAccountRequest struct {
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	Address     string                 `pathParam:"style=simple,explode=false,name=address"`
	Ledger      string                 `pathParam:"style=simple,explode=false,name=ledger"`
}

type AddMetadataToAccountResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	RawResponse   *http.Response
}
