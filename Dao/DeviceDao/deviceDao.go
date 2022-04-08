package DeviceDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/Device"
	"fmt"
	"gorm.io/gorm"
)

type DeviceDao struct {
	gorm.Model
	DeviceTypeID uint `gorm:"type:uint"`
	DeviceStatus int  `gorm:"type:int"`
}

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&DeviceDao{})
	} else {
		fmt.Println("Error happened when initializing DeviceDao and creating Table in function DeviceDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	connectDatabase()
}
func (DeviceDao) TableName() string {
	return "devices"
}
func convertDaoToDevice(dao DeviceDao) Device.Device {
	return Device.Device{
		DeviceID:     dao.ID,
		DeviceTypeID: dao.DeviceTypeID,
		DeviceStatus: dao.DeviceStatus,
	}
}
func convertDeviceToDao(device Device.Device) DeviceDao {
	return DeviceDao{
		DeviceTypeID: device.DeviceTypeID,
		DeviceStatus: device.DeviceStatus,
	}
}
func InsertDevice(device Device.Device) error {
	var deviceDao = convertDeviceToDao(device)
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&deviceDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when inserting devices in function DeviceDao.InsertDevice()")
	}
	return err
}
func UpdateDevice(device Device.Device) error {
	var deviceDao = convertDeviceToDao(device)
	deviceDao.ID = device.DeviceID
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Model(&deviceDao).Updates(deviceDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when updating devices in function DeviceDao.UpdateDevice()")
	}
	return err
}

func FindDeviceByType(deviceType uint) ([]Device.Device, error) {
	var daos []DeviceDao
	var devices []Device.Device
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where(&DeviceDao{DeviceTypeID: deviceType, DeviceStatus: 2}).Find(&daos).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding devices in function DeviceDao.FindDeviceByType()")
		fmt.Println(err)
	}
	devices = make([]Device.Device, len(daos), len(daos))
	for key := range daos {
		devices[key] = convertDaoToDevice(daos[key])
	}
	return devices, err
}
