package ReservationInfoDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/ReservationInfo"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ReservationInfoDao struct {
	gorm.Model
	ReservationID   uint   `gorm:"uint"`
	DeviceID        uint   `gorm:"uint"`
	DeviceTypeInfo  string `gorm:"type:string"`
	ReservationDay  string `gorm:"string"`
	ReservationTime int    `gorm:"int"`
}

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&ReservationInfoDao{})
	} else {
		fmt.Println("Error happened when initializing ReservationInfoDao and creating Table in function ReservationInfoDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	connectDatabase()
}
func (ReservationInfoDao) TableName() string {
	return "reservationInfo"
}
func convertInfoToDao(info ReservationInfo.ReservationInfo) ReservationInfoDao {
	return ReservationInfoDao{
		ReservationID:   info.ReservationID,
		DeviceID:        info.DeviceID,
		DeviceTypeInfo:  info.DeviceTypeInfo,
		ReservationDay:  info.ReservationDay,
		ReservationTime: info.ReservationTime,
	}
}
func convertDaoToInfo(dao ReservationInfoDao) ReservationInfo.ReservationInfo {
	return ReservationInfo.ReservationInfo{
		ID:              dao.ID,
		ReservationID:   dao.ReservationID,
		DeviceID:        dao.DeviceID,
		DeviceTypeInfo:  dao.DeviceTypeInfo,
		ReservationDay:  dao.ReservationDay,
		ReservationTime: dao.ReservationTime,
	}
}
func InsertReservationInfo(info ReservationInfo.ReservationInfo) error {
	var dao = convertInfoToDao(info)
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&dao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when inserting reservationInfo in function ReservationInfoDao.InsertReservationInfo()")
	}
	return err
}
func FindInfoByReservationID(id uint) ([]ReservationInfo.ReservationInfo, error) {
	var daos []ReservationInfoDao
	var info []ReservationInfo.ReservationInfo
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("reservation_id=?", id).Find(&daos).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when Finding reservationInfo in function ReservationInfoDao.FindInfoByReservationID()")
	} else {
		info = make([]ReservationInfo.ReservationInfo, len(daos), len(daos))
		for key := range daos {
			info[key] = convertDaoToInfo(daos[key])
		}
	}
	return info, err
}
func FindAllReservationInfo() ([]ReservationInfo.ReservationInfo, error) {
	var dao []ReservationInfoDao
	var info []ReservationInfo.ReservationInfo
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&dao).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		fmt.Println("Error happened when Finding reservationInfo in function ReservationInfoDao.FindInfoByReservationByType()")
	} else {
		info = make([]ReservationInfo.ReservationInfo, len(dao), len(dao))
		for k := range dao {
			info[k] = convertDaoToInfo(dao[k])
		}
	}
	return info, err
}
