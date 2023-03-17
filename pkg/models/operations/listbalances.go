package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ListBalancesRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type ListBalancesResponse struct {
	ContentType          string
	ListBalancesResponse *shared.ListBalancesResponse
	StatusCode           int
	RawResponse          *http.Response
}
