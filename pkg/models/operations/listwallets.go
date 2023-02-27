package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type ListWalletsQueryParams struct {
	Cursor   *string                `queryParam:"style=form,explode=true,name=cursor"`
	Metadata map[string]interface{} `queryParam:"style=deepObject,explode=true,name=metadata"`
	Name     *string                `queryParam:"style=form,explode=true,name=name"`
	PageSize *int64                 `queryParam:"style=form,explode=true,name=pageSize"`
}

type ListWalletsRequest struct {
	QueryParams ListWalletsQueryParams
}

type ListWalletsResponse struct {
	ContentType         string
	ListWalletsResponse *shared.ListWalletsResponse
	StatusCode          int
}
