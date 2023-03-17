package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ListWalletsRequest struct {
	Cursor   *string                `queryParam:"style=form,explode=true,name=cursor"`
	Metadata map[string]interface{} `queryParam:"style=deepObject,explode=true,name=metadata"`
	Name     *string                `queryParam:"style=form,explode=true,name=name"`
	PageSize *int64                 `queryParam:"style=form,explode=true,name=pageSize"`
}

type ListWalletsResponse struct {
	ContentType         string
	ListWalletsResponse *shared.ListWalletsResponse
	StatusCode          int
	RawResponse         *http.Response
}
