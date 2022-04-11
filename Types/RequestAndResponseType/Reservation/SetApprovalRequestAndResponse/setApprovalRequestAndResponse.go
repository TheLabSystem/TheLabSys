package SetApprovalRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type SetApprovalRequest struct {
	ReservationID uint    `json:"reservation_id"`
	Approval      int     `json:"approval"`
	Money         float64 `json:"money"`
}
type SetApprovalResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
