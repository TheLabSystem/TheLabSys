package GetReservationDetailsRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type GetReservationDetailsRequest struct {
	Day        string `json:"day"`
	DeviceType uint   `json:"device_type"`
}

type GetReservationDetailsResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Nums    []int  `json:"nums"`
		Message string `json:"message"`
	}
}
