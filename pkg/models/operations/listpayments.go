package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type ListPaymentsQueryParams struct {
	Cursor   *string  `queryParam:"style=form,explode=true,name=cursor"`
	PageSize *int64   `queryParam:"style=form,explode=true,name=pageSize"`
	Sort     []string `queryParam:"style=form,explode=true,name=sort"`
}

type ListPaymentsRequest struct {
	QueryParams ListPaymentsQueryParams
}

type ListPaymentsResponse struct {
	ContentType    string
	PaymentsCursor *shared.PaymentsCursor
	StatusCode     int
}
