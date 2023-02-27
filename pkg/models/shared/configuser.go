package shared

type ConfigUser struct {
	Endpoint   string   `json:"endpoint"`
	EventTypes []string `json:"eventTypes"`
	Secret     *string  `json:"secret,omitempty"`
}
