package PayBillRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type PayBillRequest struct {
	BillID uint `json:"bill_id"`
}

type PayBillResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
