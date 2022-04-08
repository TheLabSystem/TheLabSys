package GetDeviceTypeRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type ResponseInfo struct {
	DeviceTypeID uint   `json:"device_type_id"`
	DeviceInfo   string `json:"device_info"`
}

type GetDeviceTypeRequest struct {
}
type GetDeviceTypeResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message        string `json:"message"`
		DeviceTypeInfo []ResponseInfo
	}
}
