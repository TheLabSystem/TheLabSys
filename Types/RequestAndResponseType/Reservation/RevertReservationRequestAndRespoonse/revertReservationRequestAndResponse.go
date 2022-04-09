package RevertReservationRequestAndRespoonse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type RevertReservationRequest struct {
	ReservationID uint `json:"reservation_id"`
}
type RevertReservationResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
