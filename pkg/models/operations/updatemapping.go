package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateMappingRequest struct {
	Mapping shared.Mapping `request:"mediaType=application/json"`
	Ledger  string         `pathParam:"style=simple,explode=false,name=ledger"`
}

type UpdateMappingResponse struct {
	ContentType     string
	ErrorResponse   *shared.ErrorResponse
	MappingResponse *shared.MappingResponse
	StatusCode      int
	RawResponse     *http.Response
}
