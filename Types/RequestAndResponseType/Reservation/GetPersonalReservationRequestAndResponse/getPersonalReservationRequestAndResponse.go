package GetPersonalReservationRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/Reservation"
)

type GetPersonalReservationRequest struct {
}
type GetPersonalReservationResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Reservations []Reservation.Reservation `json:"reservations"`
		Message      string                    `json:"message"`
	}
}
