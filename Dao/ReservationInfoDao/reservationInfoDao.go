package ReservationInfoDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/ReservationInfo"
	"fmt"
	"gorm.io/gorm"
)

type ReservationInfoDao struct {
	gorm.Model
	ReservationID   uint   `gorm:"uint"`
	DeviceID        uint   `gorm:"uint"`
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
func convertInfoToDao(info ReservationInfo.ReservationInfo) ReservationInfoDao {
	return ReservationInfoDao{
		ReservationID:   info.ReservationID,
		DeviceID:        info.DeviceID,
		ReservationDay:  info.ReservationDay,
		ReservationTime: info.ReservationTime,
	}
}
func convertDaoToInfo(dao ReservationInfoDao) ReservationInfo.ReservationInfo {
	return ReservationInfo.ReservationInfo{
		ID:              dao.ID,
		ReservationID:   dao.ReservationID,
		DeviceID:        dao.DeviceID,
		ReservationDay:  dao.ReservationDay,
		ReservationTime: dao.ReservationTime,
	}
}
