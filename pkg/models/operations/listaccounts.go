package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type ListAccountsPathParams struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type ListAccountsBalanceOperatorEnum string

const (
	ListAccountsBalanceOperatorEnumGte ListAccountsBalanceOperatorEnum = "gte"
	ListAccountsBalanceOperatorEnumLte ListAccountsBalanceOperatorEnum = "lte"
	ListAccountsBalanceOperatorEnumGt  ListAccountsBalanceOperatorEnum = "gt"
	ListAccountsBalanceOperatorEnumLt  ListAccountsBalanceOperatorEnum = "lt"
	ListAccountsBalanceOperatorEnumE   ListAccountsBalanceOperatorEnum = "e"
	ListAccountsBalanceOperatorEnumNe  ListAccountsBalanceOperatorEnum = "ne"
)

type ListAccountsQueryParams struct {
	Address                   *string                          `queryParam:"style=form,explode=true,name=address"`
	After                     *string                          `queryParam:"style=form,explode=true,name=after"`
	Balance                   *int64                           `queryParam:"style=form,explode=true,name=balance"`
	BalanceOperator           *ListAccountsBalanceOperatorEnum `queryParam:"style=form,explode=true,name=balanceOperator"`
	BalanceOperatorDeprecated *ListAccountsBalanceOperatorEnum `queryParam:"style=form,explode=true,name=balance_operator"`
	Cursor                    *string                          `queryParam:"style=form,explode=true,name=cursor"`
	Metadata                  map[string]interface{}           `queryParam:"style=deepObject,explode=true,name=metadata"`
	PageSize                  *int64                           `queryParam:"style=form,explode=true,name=pageSize"`
	PageSizeDeprecated        *int64                           `queryParam:"style=form,explode=true,name=page_size"`
	PaginationToken           *string                          `queryParam:"style=form,explode=true,name=pagination_token"`
}

type ListAccountsRequest struct {
	PathParams  ListAccountsPathParams
	QueryParams ListAccountsQueryParams
}

type ListAccountsResponse struct {
	AccountsCursorResponse *shared.AccountsCursorResponse
	ContentType            string
	ErrorResponse          *shared.ErrorResponse
	StatusCode             int
}
