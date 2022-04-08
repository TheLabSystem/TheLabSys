package DeviceRecordDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/DeviceRecord"
	"fmt"
	"gorm.io/gorm"
)

type DeviceRecordDao struct {
	gorm.Model
	DeviceOperationID uint `gorm:"type:uint"`
	DeviceID          uint `gorm:"type:uint"`
	OperationType     int  `gorm:"type:int"`
}

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&DeviceRecordDao{})
	} else {
		fmt.Println("Error happened when initializing DeviceRecordDao and creating Table in function DeviceRecordDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	connectDatabase()
}
func (DeviceRecordDao) TableName() string {
	return "deviceRecords"
}
func convertDaoToDeviceRecord(dao DeviceRecordDao) DeviceRecord.DeviceRecord {
	return DeviceRecord.DeviceRecord{
		ID:                dao.ID,
		DeviceOperationID: dao.DeviceOperationID,
		DeviceID:          dao.DeviceID,
		OperationType:     dao.OperationType,
	}
}
func convertDeviceRecordToDao(dr DeviceRecord.DeviceRecord) DeviceRecordDao {
	return DeviceRecordDao{
		DeviceOperationID: dr.DeviceOperationID,
		DeviceID:          dr.DeviceID,
		OperationType:     dr.OperationType,
	}
}
func InsertDeviceRecord(dr DeviceRecord.DeviceRecord) error {
	var drDao = convertDeviceRecordToDao(dr)
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&drDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when inserting DeviceRecords in function DeviceRecordDao.InsertDeviceRecord()")
	}
	return err
}
