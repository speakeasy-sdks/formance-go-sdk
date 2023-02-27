package shared

type ConnectorsResponseData struct {
	Enabled  *bool          `json:"enabled,omitempty"`
	Provider *ConnectorEnum `json:"provider,omitempty"`
}

type ConnectorsResponse struct {
	Data []ConnectorsResponseData `json:"data"`
}
