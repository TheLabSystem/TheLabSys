package DeviceService

import (
	"TheLabSystem/Config/UserPermissionDecide"
	"TheLabSystem/Dao/DeviceDao"
	"TheLabSystem/Dao/DeviceOperationDao"
	"TheLabSystem/Dao/DeviceRecordDao"
	"TheLabSystem/Dao/DeviceTypeInfoDao"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/RequestAndResponseType/Device/GetDeviceTypeRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Device/GetDevicesRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/Device"
	"TheLabSystem/Types/ServiceType/DeviceOperation"
	"TheLabSystem/Types/ServiceType/DeviceRecord"
	"TheLabSystem/Types/ServiceType/DeviceTypeInfo"
	"errors"
	"gorm.io/gorm"
)

//UpdateDevice(); Done
//DeleteDevice();
//GetDeviceType(); Done
//AddDevice(); Done
//AllDone

type DeviceService struct {
}

func (service DeviceService) AddDevice(username string, deviceID uint, deviceTypeID uint, deviceInfo string, money float64) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username == "" {
		return ErrNo.LoginRequired
	}
	if !UserPermissionDecide.AddDevice(user.UserType) {
		return ErrNo.PermDenied
	}
	if err := DeviceDao.InsertDevice(Device.Device{
		DeviceID:     deviceID,
		DeviceTypeID: deviceTypeID,
		DeviceStatus: 2,
	}); err != nil {
		return ErrNo.UnknownError
	}
	operation, err := DeviceOperationDao.InsertDeviceOperation(DeviceOperation.DeviceOperation{
		OperatorID: user.UserID,
	})
	if err != nil {
		return ErrNo.UnknownError
	}
	if err := DeviceRecordDao.InsertDeviceRecord(DeviceRecord.DeviceRecord{
		DeviceOperationID: operation.ID,
		DeviceID:          deviceID,
		OperationType:     1,
	}); err != nil {
		return ErrNo.UnknownError
	}
	if _, err := DeviceTypeInfoDao.FindDeviceTypeInfoByDeviceTypeID(deviceTypeID); errors.Is(err, gorm.ErrRecordNotFound) {
		err := DeviceTypeInfoDao.InsertDeviceTypeInfo(DeviceTypeInfo.DeviceTypeInfo{
			DeviceTypeID: deviceTypeID,
			DeviceInfo:   deviceInfo,
			Money:        money,
		})
		if err != nil {
			return ErrNo.UnknownError
		}
	}
	return ErrNo.OK
}
func (service DeviceService) GetDeviceType(username string) (ErrNo.ErrNo, []GetDeviceTypeRequestAndResponse.ResponseInfo) {
	var responseInfo []GetDeviceTypeRequestAndResponse.ResponseInfo
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError, responseInfo
	}
	if user.Username == "" {
		return ErrNo.LoginRequired, responseInfo
	}
	if !UserPermissionDecide.AddDevice(user.UserType) {
		return ErrNo.PermDenied, responseInfo
	}
	if deviceTypeInfo, err := DeviceTypeInfoDao.FindAllDeviceTypeInfo(); err != nil {
		return ErrNo.UnknownError, responseInfo
	} else {
		responseInfo = make([]GetDeviceTypeRequestAndResponse.ResponseInfo, len(deviceTypeInfo), len(deviceTypeInfo))
		for key := range deviceTypeInfo {
			responseInfo[key].DeviceTypeID = deviceTypeInfo[key].DeviceTypeID
			responseInfo[key].DeviceInfo = deviceTypeInfo[key].DeviceInfo
		}
	}
	return ErrNo.OK, responseInfo
}
func (service DeviceService) UpdateDevice(username string, deviceID uint, deviceTypeID uint, deviceStatus int) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username == "" {
		return ErrNo.LoginRequired
	}
	if !UserPermissionDecide.UpdateDevice(user.UserType) {
		return ErrNo.PermDenied
	}
	if err := DeviceDao.UpdateDevice(Device.Device{
		DeviceID:     deviceID,
		DeviceTypeID: deviceTypeID,
		DeviceStatus: deviceStatus,
	}); err != nil {
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
func (service DeviceService) DeleteDevice(username string, deviceID uint) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username == "" {
		return ErrNo.LoginRequired
	}
	if !UserPermissionDecide.AddDevice(user.UserType) {
		return ErrNo.PermDenied
	}
	if err := DeviceDao.DeleteDevice(deviceID); err != nil {
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
func (service DeviceService) GetDevices(username string) ([]GetDevicesRequestAndResponse.DeviceTypeAndIDs, ErrNo.ErrNo) {
	var res []GetDevicesRequestAndResponse.DeviceTypeAndIDs
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return res, ErrNo.UnknownError
	}
	if user.Username == "" {
		return res, ErrNo.LoginRequired
	}
	var deviceTypes []DeviceTypeInfo.DeviceTypeInfo
	deviceTypes, err = DeviceTypeInfoDao.FindAllDeviceTypeInfo()
	if err != nil {
		return res, ErrNo.UnknownError
	}
	res = make([]GetDevicesRequestAndResponse.DeviceTypeAndIDs, len(deviceTypes), len(deviceTypes))
	for k := range deviceTypes {
		var res1 = GetDevicesRequestAndResponse.DeviceTypeAndIDs{
			TypeInfo: deviceTypes[k],
		}
		res1.Devices, err = DeviceDao.FindDeviceByTypeID(deviceTypes[k].DeviceTypeID)
		if err != nil {
			return res, ErrNo.UnknownError
		}
		res[k] = res1
	}
	return res, ErrNo.OK
}
