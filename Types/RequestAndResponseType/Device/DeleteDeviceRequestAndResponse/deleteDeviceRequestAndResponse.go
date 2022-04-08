package DeleteDeviceRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
)

type DeleteDeviceRequest struct {
	DeviceID uint `json:"device_id"`
}
type DeleteDeviceResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
