package GetReservationInfoByReservationIDRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/Reservation"
	"TheLabSystem/Types/ServiceType/ReservationInfo"
)

type GetReservationInfoByReservationIDRequest struct {
	ReservationID uint `json:"reservation_id"`
}
type GetReservationInfoByReservationIDResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Reservation      Reservation.Reservation           `json:"reservation"`
		ReservationInfos []ReservationInfo.ReservationInfo `json:"reservationInfos"`
		Message          string                            `json:"message"`
	}
}
