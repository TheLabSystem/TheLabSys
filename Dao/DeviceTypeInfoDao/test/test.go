package main

import (
	"TheLabSystem/Dao/DeviceTypeInfoDao"
	"TheLabSystem/Types/ServiceType/DeviceTypeInfo"
)

func main() {
	info := DeviceTypeInfo.DeviceTypeInfo{
		DeviceTypeID: 1,
		DeviceInfo:   "Hello World",
	}
	DeviceTypeInfoDao.InsertDeviceTypeInfo(info)
}
