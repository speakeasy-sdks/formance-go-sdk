package shared

type WalletsVolume struct {
	Balance int64 `json:"balance"`
	Input   int64 `json:"input"`
	Output  int64 `json:"output"`
}
