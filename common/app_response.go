package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func SuccessRepsone(data, paging, filter interface{}) *successRes {
	return &successRes{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}
func SimpleSuccessRepsone(data interface{}) successRes {
	return successRes{
		data,
		nil,
		nil,
	}
}
