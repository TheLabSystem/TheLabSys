package SubmitReservationRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type SubmitReservationRequest struct {
	DeviceType uint   `json:"deviceType"`
	Day        string `json:"day"`
	Time       int    `json:"time"`
	Num        int    `json:"num"`
}

type SubmitReservationResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
