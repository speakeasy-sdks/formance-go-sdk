// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ResponseCursorData struct {
}

type ResponseCursorTotal struct {
	Relation *string `json:"relation,omitempty"`
	Value    *int64  `json:"value,omitempty"`
}

func (o *ResponseCursorTotal) GetRelation() *string {
	if o == nil {
		return nil
	}
	return o.Relation
}

func (o *ResponseCursorTotal) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

type ResponseCursor struct {
	Data     []ResponseCursorData `json:"data,omitempty"`
	HasMore  *bool                `json:"hasMore,omitempty"`
	Next     *string              `json:"next,omitempty"`
	PageSize *int64               `json:"pageSize,omitempty"`
	Previous *string              `json:"previous,omitempty"`
	Total    *ResponseCursorTotal `json:"total,omitempty"`
}

func (o *ResponseCursor) GetData() []ResponseCursorData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ResponseCursor) GetHasMore() *bool {
	if o == nil {
		return nil
	}
	return o.HasMore
}

func (o *ResponseCursor) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ResponseCursor) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ResponseCursor) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ResponseCursor) GetTotal() *ResponseCursorTotal {
	if o == nil {
		return nil
	}
	return o.Total
}

// ResponseData - The payload
type ResponseData struct {
}

type Response struct {
	Cursor *ResponseCursor `json:"cursor,omitempty"`
	// The payload
	Data *ResponseData `json:"data,omitempty"`
}

func (o *Response) GetCursor() *ResponseCursor {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *Response) GetData() *ResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}
