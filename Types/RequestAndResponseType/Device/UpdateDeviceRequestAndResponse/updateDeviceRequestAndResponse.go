package UpdateDeviceRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type UpdateDeviceRequest struct {
	DeviceID     uint `json:"device_id"`
	DeviceTypeID uint `json:"device_type_id"`
	DeviceStatus int  `json:"device_status"`
}
type UpdateDeviceResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
