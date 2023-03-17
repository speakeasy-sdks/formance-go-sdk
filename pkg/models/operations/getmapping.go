package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetMappingRequest struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type GetMappingResponse struct {
	ContentType     string
	ErrorResponse   *shared.ErrorResponse
	MappingResponse *shared.MappingResponse
	StatusCode      int
	RawResponse     *http.Response
}
