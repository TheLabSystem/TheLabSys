package DeviceTypeInfoDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/DeviceTypeInfo"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type DeviceTypeInfoDao struct {
	gorm.Model
	DeviceTypeID uint    `gorm:"type:uint"`
	DeviceInfo   string  `gorm:"type:string"`
	Money        float64 `gorm:"type:float"`
}

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&DeviceTypeInfoDao{})
	} else {
		fmt.Println("Error happened when initializing NoticeDao and creating Table in function NoticeDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	connectDatabase()
}
func (DeviceTypeInfoDao) TableName() string {
	return "deviceTypeInfo"
}
func convertDaoToDeviceTypeInfo(dao DeviceTypeInfoDao) DeviceTypeInfo.DeviceTypeInfo {
	return DeviceTypeInfo.DeviceTypeInfo{
		DeviceTypeID: dao.DeviceTypeID,
		DeviceInfo:   dao.DeviceInfo,
		Money:        dao.Money,
	}
}
func convertDeviceTypeInfoToDao(info DeviceTypeInfo.DeviceTypeInfo) DeviceTypeInfoDao {
	return DeviceTypeInfoDao{
		DeviceInfo:   info.DeviceInfo,
		DeviceTypeID: info.DeviceTypeID,
		Money:        info.Money,
	}
}
func FindDeviceTypeInfoByDeviceTypeID(id uint) (DeviceTypeInfo.DeviceTypeInfo, error) {
	var dao DeviceTypeInfoDao
	var info DeviceTypeInfo.DeviceTypeInfo
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where(&DeviceTypeInfoDao{DeviceTypeID: id}).First(&dao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	//fmt.Println(err)
	if err == nil {
		info = convertDaoToDeviceTypeInfo(dao)
	}
	return info, err
}
func InsertDeviceTypeInfo(info DeviceTypeInfo.DeviceTypeInfo) error {
	var dao = convertDeviceTypeInfoToDao(info)
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&dao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when inserting deviceTypeInfo in function DeviceTypeInfoDao.InsertDeviceTypeInfo()")
	}
	return err
}
func FindAllDeviceTypeInfo() ([]DeviceTypeInfo.DeviceTypeInfo, error) {
	var dao []DeviceTypeInfoDao
	var info []DeviceTypeInfo.DeviceTypeInfo
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Find(&dao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	} else if err != nil {
		fmt.Println("Error happened when finding deviceTypeInfo in function DeviceTypeInfoDao.FindAllDeviceTypeInfo()")
		fmt.Println(err)
	} else {
		info = make([]DeviceTypeInfo.DeviceTypeInfo, len(dao), len(dao))
		for key := range dao {
			info[key] = convertDaoToDeviceTypeInfo(dao[key])
		}
	}
	return info, err
}
