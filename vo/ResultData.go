package vo

type ResultData[T any, V any] struct {
	Code    int32  `json:"Code"`
	Message string `json:"Message"`
	ExtData V      `json:"ExtData"`
	Data    T      `json:"Data"`
}
