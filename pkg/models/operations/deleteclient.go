package operations

type DeleteClientPathParams struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

type DeleteClientRequest struct {
	PathParams DeleteClientPathParams
}

type DeleteClientResponse struct {
	ContentType string
	StatusCode  int
}
