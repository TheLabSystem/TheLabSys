package GetBillRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/Bill"
)

type GetBillRequest struct {
}

type GetBillResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string      `json:"message"`
		Bills   []Bill.Bill `json:"bills"`
	}
}
