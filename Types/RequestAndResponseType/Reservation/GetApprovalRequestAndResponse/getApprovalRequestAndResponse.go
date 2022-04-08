package GetApprovalRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/Reservation"
)

type GetApprovalRequest struct {
	Status int `json:"status"`
}
type GetApprovalResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message  string `json:"message"`
		Approval []Reservation.Reservation
	}
}
