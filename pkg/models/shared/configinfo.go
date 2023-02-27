package shared

type ConfigInfo struct {
	Config  Config `json:"config"`
	Server  string `json:"server"`
	Version string `json:"version"`
}
