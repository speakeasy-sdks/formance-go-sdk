package shared

type ConnectorsConfigsResponseDataConnectorKey struct {
	DataType string `json:"dataType"`
	Required bool   `json:"required"`
}

type ConnectorsConfigsResponseDataConnector struct {
	Key ConnectorsConfigsResponseDataConnectorKey `json:"key"`
}

type ConnectorsConfigsResponseData struct {
	Connector ConnectorsConfigsResponseDataConnector `json:"connector"`
}

type ConnectorsConfigsResponse struct {
	Data ConnectorsConfigsResponseData `json:"data"`
}
