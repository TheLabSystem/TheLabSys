package GetApprovalRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/Reservation"
	"TheLabSystem/Types/ServiceType/User"
)

type Approval struct {
	ReservationRes Reservation.Reservation
	UserRes        User.User
}
type GetApprovalRequest struct {
	Status int `json:"status"`
}
type GetApprovalResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message     string `json:"message"`
		ApprovalRes []Approval
	}
}
