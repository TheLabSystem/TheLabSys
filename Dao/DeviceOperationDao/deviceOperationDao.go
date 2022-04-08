package DeviceOperationDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/DeviceOperation"
	"fmt"
	"gorm.io/gorm"
)

type DeviceOperationDao struct {
	gorm.Model
	OperatorID uint `gorm:"type:uint"`
}

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&DeviceOperationDao{})
	} else {
		fmt.Println("Error happened when initializing DeviceOperationDao and creating Table in function DeviceOperationDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	connectDatabase()
}
func (DeviceOperationDao) TableName() string {
	return "deviceOperations"
}
func convertDaoToDeviceOperation(dao DeviceOperationDao) DeviceOperation.DeviceOperation {
	return DeviceOperation.DeviceOperation{
		ID:         dao.ID,
		OperatorID: dao.OperatorID,
	}
}
func convertDeviceOperationToDao(operation DeviceOperation.DeviceOperation) DeviceOperationDao {
	return DeviceOperationDao{
		OperatorID: operation.OperatorID,
	}
}
func InsertDeviceOperation(operation DeviceOperation.DeviceOperation) error {
	dao := convertDeviceOperationToDao(operation)
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&dao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when inserting deviceOperation in function DeviceOperationDao.InsertDeviceOperation()")
	}
	return err
}
