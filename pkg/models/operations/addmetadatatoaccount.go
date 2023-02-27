package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type AddMetadataToAccountPathParams struct {
	Address string `pathParam:"style=simple,explode=false,name=address"`
	Ledger  string `pathParam:"style=simple,explode=false,name=ledger"`
}

type AddMetadataToAccountRequest struct {
	PathParams AddMetadataToAccountPathParams
	Request    map[string]interface{} `request:"mediaType=application/json"`
}

type AddMetadataToAccountResponse struct {
	ContentType   string
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
}
