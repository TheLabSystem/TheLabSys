package AddDeviceRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
)

type AddDeviceRequest struct {
	DeviceID     uint   `json:"device_id"`
	DeviceTypeID uint   `json:"device_type_id"`
	DeviceInfo   string `json:"device_info"`
}

type AddDeviceResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"Message"`
	}
}
