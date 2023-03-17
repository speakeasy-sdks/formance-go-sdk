package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetHoldsRequest struct {
	Cursor   *string                `queryParam:"style=form,explode=true,name=cursor"`
	Metadata map[string]interface{} `queryParam:"style=deepObject,explode=true,name=metadata"`
	PageSize *int64                 `queryParam:"style=form,explode=true,name=pageSize"`
	WalletID *string                `queryParam:"style=form,explode=true,name=walletID"`
}

type GetHoldsResponse struct {
	ContentType          string
	GetHoldsResponse     *shared.GetHoldsResponse
	StatusCode           int
	RawResponse          *http.Response
	WalletsErrorResponse *shared.WalletsErrorResponse
}
