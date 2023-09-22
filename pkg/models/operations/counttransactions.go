// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/utils"
	"net/http"
	"time"
)

// CountTransactionsMetadata - Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below.
type CountTransactionsMetadata struct {
}

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
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	EndTimeDeprecated *time.Time `queryParam:"style=form,explode=true,name=end_time"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below.
	Metadata *CountTransactionsMetadata `queryParam:"style=deepObject,explode=true,name=metadata"`
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
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	StartTimeDeprecated *time.Time `queryParam:"style=form,explode=true,name=start_time"`
}

func (c CountTransactionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CountTransactionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CountTransactionsRequest) GetAccount() *string {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *CountTransactionsRequest) GetDestination() *string {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *CountTransactionsRequest) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *CountTransactionsRequest) GetEndTimeDeprecated() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTimeDeprecated
}

func (o *CountTransactionsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *CountTransactionsRequest) GetMetadata() *CountTransactionsMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *CountTransactionsRequest) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *CountTransactionsRequest) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *CountTransactionsRequest) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *CountTransactionsRequest) GetStartTimeDeprecated() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTimeDeprecated
}

type CountTransactionsResponse struct {
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
	StatusCode    int
	RawResponse   *http.Response
}

func (o *CountTransactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CountTransactionsResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *CountTransactionsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CountTransactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CountTransactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
