package GetDevicesRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/Device"
	"TheLabSystem/Types/ServiceType/DeviceTypeInfo"
)

type DeviceTypeAndIDs struct {
	TypeInfo DeviceTypeInfo.DeviceTypeInfo `json:"info"`
	Devices  []Device.Device               `json:"devices"`
}
type GetDevicesRequest struct {
}
type GetDevicesResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string             `json:"message"`
		Devices []DeviceTypeAndIDs `json:"devices"`
	}
}
