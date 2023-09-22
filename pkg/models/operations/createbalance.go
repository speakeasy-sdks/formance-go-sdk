// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateBalanceRequest struct {
	CreateBalanceRequest *shared.CreateBalanceRequest `request:"mediaType=application/json"`
	ID                   string                       `pathParam:"style=simple,explode=false,name=id"`
}

func (o *CreateBalanceRequest) GetCreateBalanceRequest() *shared.CreateBalanceRequest {
	if o == nil {
		return nil
	}
	return o.CreateBalanceRequest
}

func (o *CreateBalanceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type CreateBalanceResponse struct {
	ContentType string
	// Created balance
	CreateBalanceResponse *shared.CreateBalanceResponse
	StatusCode            int
	RawResponse           *http.Response
	// Error
	WalletsErrorResponse *shared.WalletsErrorResponse
}

func (o *CreateBalanceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateBalanceResponse) GetCreateBalanceResponse() *shared.CreateBalanceResponse {
	if o == nil {
		return nil
	}
	return o.CreateBalanceResponse
}

func (o *CreateBalanceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateBalanceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateBalanceResponse) GetWalletsErrorResponse() *shared.WalletsErrorResponse {
	if o == nil {
		return nil
	}
	return o.WalletsErrorResponse
}
