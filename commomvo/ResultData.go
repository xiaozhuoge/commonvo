package commomvo

type ResultData[T any, V any] struct {
	Code    int32  `json:"Code"`
	Message string `json:"Message"`
	ExtData V      `json:"ExtData"`
	Data    T      `json:"Data"`
}

type ResultDataWithoutExtParameters[T any] struct {
	Code    int32  `json:"Code"`
	Message string `json:"Message"`
	Data    T      `json:"Data"`
}

type ResultDataWithTwoExtParameters[T any, V any, B any] struct {
	Code       int32  `json:"Code"`
	Message    string `json:"Message"`
	ExtData    V      `json:"ExtData"`
	ExtDataTwo B      `json:"ExtDataTwo"`
	Data       T      `json:"Data"`
}

type ResultDataWithThreeExtParameters[T any, V any, B any, C any] struct {
	Code         int32  `json:"Code"`
	Message      string `json:"Message"`
	ExtData      V      `json:"ExtData"`
	ExtDataTwo   B      `json:"ExtDataTwo"`
	ExtDataThree C      `json:"ExtDataThree"`
	Data         T      `json:"Data"`
}

type ResultDataWithFourExtParameters[T any, V any, B any, C any, D any] struct {
	Code         int32  `json:"Code"`
	Message      string `json:"Message"`
	ExtData      V      `json:"ExtData"`
	ExtDataTwo   B      `json:"ExtDataTwo"`
	ExtDataThree C      `json:"ExtDataThree"`
	ExtDataFour  D      `json:"ExtDataFour"`
	Data         T      `json:"Data"`
}

type ResultDataWithFiveExtParameters[T any, V any, B any, C any, D any, E any] struct {
	Code         int32  `json:"Code"`
	Message      string `json:"Message"`
	ExtData      V      `json:"ExtData"`
	ExtDataTwo   B      `json:"ExtDataTwo"`
	ExtDataThree C      `json:"ExtDataThree"`
	ExtDataFour  D      `json:"ExtDataFour"`
	ExtDataFive  E      `json:"ExtDataFive"`
	Data         T      `json:"Data"`
}

type PageBean[T any] struct {
	TotalCount  int64 `json:"TotalCount"`
	CurrentData []T   `json:"CurrentData"`
}
