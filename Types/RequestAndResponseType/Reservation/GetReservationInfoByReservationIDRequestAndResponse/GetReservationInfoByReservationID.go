package GetReservationInfoByReservationIDRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/ReservationInfo"
)

type GetReservationInfoByReservationIDRequest struct {
	ReservationID uint `json:"reservation_id"`
}
type GetReservationInfoByReservationIDResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		ReservationInfo ReservationInfo.ReservationInfo `json:"reservationInfo"`
		Message         string                          `json:"message"`
	}
}
