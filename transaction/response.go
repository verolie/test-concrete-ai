package transaction

import "example/transaction/model"

func ResponseDataDetail(data interface{}) model.DataStatusResponse {
	resp := model.DataStatusResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}
	return resp
}