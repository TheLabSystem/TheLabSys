package ReservationRecordDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/ReservationRecord"
	"fmt"
	"gorm.io/gorm"
)

type ReservationRecordDao struct {
	gorm.Model
	ReservationID uint   `gorm:"type:uint"`
	OperatorID    uint   `gorm:"type:uint"`
	OperationType int    `gorm:"type:integer"`
	OperatingDay  string `gorm:"type:string"`
}

// 申请:operationType=1
// 拒绝：operationType=2
// 同意：operationType=3 老师-》学生 5，
// 取消：operationType=4

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&ReservationRecordDao{})
	} else {
		fmt.Println("Error happened when initializing ReservationRecordDao and creating Table in function ReservationRecordDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	connectDatabase()
}
func (ReservationRecordDao) TableName() string {
	return "reservationRecord"
}
func convertDaoToReservationRecord(dao ReservationRecordDao) ReservationRecord.ReservationRecord {
	return ReservationRecord.ReservationRecord{
		ID:            dao.ID,
		ReservationID: dao.ReservationID,
		OperatorID:    dao.OperatorID,
		OperationType: dao.OperationType,
		OperatingDay:  dao.OperatingDay,
	}
}
func convertReservationRecordToDao(rr ReservationRecord.ReservationRecord) ReservationRecordDao {
	return ReservationRecordDao{
		ReservationID: rr.ReservationID,
		OperatorID:    rr.OperatorID,
		OperationType: rr.OperationType,
		OperatingDay:  rr.OperatingDay,
	}
}
func InsertReservationRecord(rr ReservationRecord.ReservationRecord) error {
	var rrDao = convertReservationRecordToDao(rr)
	rrDao.OperatingDay = rrDao.CreatedAt.Format("2006-04-02")
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&rrDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when inserting ReservationRecord in function ReservationRecordDao.InsertReservationRecord()")
	}
	return err
}

func FindReservationRecordByReservationID(id uint) ([]ReservationRecord.ReservationRecord, error) {
	var rrDao []ReservationRecordDao
	var rr []ReservationRecord.ReservationRecord
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("reservation_id=?", id).First(&rrDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding ReservationRecords in function ReservationRecordDao.FindReservationRecordByReservationID()")
	} else {
		rr = make([]ReservationRecord.ReservationRecord, len(rrDao), len(rrDao))
		for key := range rrDao {
			rr[key] = convertDaoToReservationRecord(rrDao[key])
		}
	}
	return rr, err
}
func FindReservationRecordByOperatorID(id uint) (ReservationRecord.ReservationRecord, error) {
	var rrDao ReservationRecordDao
	var rr ReservationRecord.ReservationRecord
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("operator_id=?", id).First(&rrDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding ReservationRecords in function ReservationRecordDao.FindReservationRecordByOperatorID()")
	} else {
		rr = convertDaoToReservationRecord(rrDao)
	}
	return rr, err
}
func FindReservationRecordByOperatingDay(day string) (ReservationRecord.ReservationRecord, error) {
	var rrDao ReservationRecordDao
	var rr ReservationRecord.ReservationRecord
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("operating_day=?", day).First(&rrDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding ReservationRecords in function ReservationRecordDao.FindReservationRecordByOperatingDay()")
	} else {
		rr = convertDaoToReservationRecord(rrDao)
	}
	return rr, err
}
