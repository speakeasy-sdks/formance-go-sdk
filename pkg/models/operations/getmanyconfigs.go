package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type GetManyConfigsQueryParams struct {
	Endpoint *string `queryParam:"style=form,explode=true,name=endpoint"`
	ID       *string `queryParam:"style=form,explode=true,name=id"`
}

type GetManyConfigsRequest struct {
	QueryParams GetManyConfigsQueryParams
}

type GetManyConfigsResponse struct {
	ConfigsResponse *shared.ConfigsResponse
	ContentType     string
	StatusCode      int
}
