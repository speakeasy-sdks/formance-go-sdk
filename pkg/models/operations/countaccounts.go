package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CountAccountsRequest struct {
	Address  *string                `queryParam:"style=form,explode=true,name=address"`
	Ledger   string                 `pathParam:"style=simple,explode=false,name=ledger"`
	Metadata map[string]interface{} `queryParam:"style=deepObject,explode=true,name=metadata"`
}

type CountAccountsResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
	StatusCode    int
	RawResponse   *http.Response
}
