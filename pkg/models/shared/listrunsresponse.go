// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ListRunsResponseCursor struct {
	Data     []WorkflowOccurrence `json:"data"`
	HasMore  *bool                `json:"hasMore,omitempty"`
	Next     *string              `json:"next,omitempty"`
	PageSize int64                `json:"pageSize"`
	Previous *string              `json:"previous,omitempty"`
}

func (o *ListRunsResponseCursor) GetData() []WorkflowOccurrence {
	if o == nil {
		return []WorkflowOccurrence{}
	}
	return o.Data
}

func (o *ListRunsResponseCursor) GetHasMore() *bool {
	if o == nil {
		return nil
	}
	return o.HasMore
}

func (o *ListRunsResponseCursor) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListRunsResponseCursor) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *ListRunsResponseCursor) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

type ListRunsResponse struct {
	Cursor ListRunsResponseCursor `json:"cursor"`
}

func (o *ListRunsResponse) GetCursor() ListRunsResponseCursor {
	if o == nil {
		return ListRunsResponseCursor{}
	}
	return o.Cursor
}
