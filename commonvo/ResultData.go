package commonvo

type ResultStatus struct {
	Code    int32  `json:"Code"`
	Message string `json:"Message"`
}

type ResultData[T any] struct {
	Code    int32  `json:"Code"`
	Message string `json:"Message"`
	Data    T      `json:"Data"`
}

type ResultDataWithOneExtParameter[T any, V any] struct {
	Code    int32  `json:"Code"`
	Message string `json:"Message"`
	ExtData V      `json:"ExtData"`
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

type ResultDataWithSixExtParameters[T any, V any, B any, C any, D any, E any, F any] struct {
	Code         int32  `json:"Code"`
	Message      string `json:"Message"`
	ExtData      V      `json:"ExtData"`
	ExtDataTwo   B      `json:"ExtDataTwo"`
	ExtDataThree C      `json:"ExtDataThree"`
	ExtDataFour  D      `json:"ExtDataFour"`
	ExtDataFive  E      `json:"ExtDataFive"`
	ExtDataSix   F      `json:"ExtDataSix"`
	Data         T      `json:"Data"`
}

type ResultDataWithSevenExtParameters[T any, V any, B any, C any, D any, E any, F any, G any] struct {
	Code         int32  `json:"Code"`
	Message      string `json:"Message"`
	ExtData      V      `json:"ExtData"`
	ExtDataTwo   B      `json:"ExtDataTwo"`
	ExtDataThree C      `json:"ExtDataThree"`
	ExtDataFour  D      `json:"ExtDataFour"`
	ExtDataFive  E      `json:"ExtDataFive"`
	ExtDataSix   F      `json:"ExtDataSix"`
	ExtDataSeven G      `json:"ExtDataSeven"`
	Data         T      `json:"Data"`
}

type ResultDataWithEightExtParameters[T any, V any, B any, C any, D any, E any, F any, G any, H any] struct {
	Code         int32  `json:"Code"`
	Message      string `json:"Message"`
	ExtData      V      `json:"ExtData"`
	ExtDataTwo   B      `json:"ExtDataTwo"`
	ExtDataThree C      `json:"ExtDataThree"`
	ExtDataFour  D      `json:"ExtDataFour"`
	ExtDataFive  E      `json:"ExtDataFive"`
	ExtDataSix   F      `json:"ExtDataSix"`
	ExtDataSeven G      `json:"ExtDataSeven"`
	ExtDataEight H      `json:"ExtDataEight"`
	Data         T      `json:"Data"`
}

type ResultDataWithNineExtParameters[T any, V any, B any, C any, D any, E any, F any, G any, H any, I any] struct {
	Code         int32  `json:"Code"`
	Message      string `json:"Message"`
	ExtData      V      `json:"ExtData"`
	ExtDataTwo   B      `json:"ExtDataTwo"`
	ExtDataThree C      `json:"ExtDataThree"`
	ExtDataFour  D      `json:"ExtDataFour"`
	ExtDataFive  E      `json:"ExtDataFive"`
	ExtDataSix   F      `json:"ExtDataSix"`
	ExtDataSeven G      `json:"ExtDataSeven"`
	ExtDataEight H      `json:"ExtDataEight"`
	ExtDataNine  I      `json:"ExtDataNine"`
	Data         T      `json:"Data"`
}

type ResultDataWithTenExtParameters[T any, V any, B any, C any, D any, E any, F any, G any, H any, I any, J any] struct {
	Code         int32  `json:"Code"`
	Message      string `json:"Message"`
	ExtData      V      `json:"ExtData"`
	ExtDataTwo   B      `json:"ExtDataTwo"`
	ExtDataThree C      `json:"ExtDataThree"`
	ExtDataFour  D      `json:"ExtDataFour"`
	ExtDataFive  E      `json:"ExtDataFive"`
	ExtDataSix   F      `json:"ExtDataSix"`
	ExtDataSeven G      `json:"ExtDataSeven"`
	ExtDataEight H      `json:"ExtDataEight"`
	ExtDataNine  I      `json:"ExtDataNine"`
	ExtDataTen   J      `json:"ExtDataTen"`
	Data         T      `json:"Data"`
}

type PageBean[T any] struct {
	TotalCount  int64 `json:"TotalCount"`
	CurrentData []T   `json:"CurrentData"`
}
