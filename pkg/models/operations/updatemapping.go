package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateMappingPathParams struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type UpdateMappingRequest struct {
	PathParams UpdateMappingPathParams
	Request    shared.Mapping `request:"mediaType=application/json"`
}

type UpdateMappingResponse struct {
	ContentType     string
	ErrorResponse   *shared.ErrorResponse
	MappingResponse *shared.MappingResponse
	StatusCode      int
	RawResponse     *http.Response
}
