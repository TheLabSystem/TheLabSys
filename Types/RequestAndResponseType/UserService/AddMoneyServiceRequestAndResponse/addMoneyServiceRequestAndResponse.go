package AddMoneyServiceRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type AddMoneyServiceRequest struct {
	Money float64 `json:"money"`
}
type AddMoneyServiceResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
