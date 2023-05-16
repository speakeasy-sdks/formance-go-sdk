// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
	"time"
)

type CountTransactionsRequest struct {
	// Filter transactions with postings involving given account, either as source or destination (regular expression placed between ^ and $).
	Account *string `queryParam:"style=form,explode=true,name=account"`
	// Filter transactions with postings involving given account at destination (regular expression placed between ^ and $).
	Destination *string `queryParam:"style=form,explode=true,name=destination"`
	// Filter transactions that occurred before this timestamp.
	// The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
	//
	EndTime *time.Time `queryParam:"style=form,explode=true,name=endTime"`
	// Filter transactions that occurred before this timestamp.
	// The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
	// Deprecated, please use `endTime` instead.
	//
	//
	// Deprecated: this field will be removed in a future release, please migrate away from it as soon as possible.
	EndTimeDeprecated *time.Time `queryParam:"style=form,explode=true,name=end_time"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below.
	Metadata map[string]interface{} `queryParam:"style=deepObject,explode=true,name=metadata"`
	// Filter transactions by reference field.
	Reference *string `queryParam:"style=form,explode=true,name=reference"`
	// Filter transactions with postings involving given account at source (regular expression placed between ^ and $).
	Source *string `queryParam:"style=form,explode=true,name=source"`
	// Filter transactions that occurred after this timestamp.
	// The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
	//
	StartTime *time.Time `queryParam:"style=form,explode=true,name=startTime"`
	// Filter transactions that occurred after this timestamp.
	// The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
	// Deprecated, please use `startTime` instead.
	//
	//
	// Deprecated: this field will be removed in a future release, please migrate away from it as soon as possible.
	StartTimeDeprecated *time.Time `queryParam:"style=form,explode=true,name=start_time"`
}

type CountTransactionsResponse struct {
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
	StatusCode    int
	RawResponse   *http.Response
}
