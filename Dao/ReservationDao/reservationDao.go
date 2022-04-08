package ReservationDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"fmt"
	"gorm.io/gorm"
)

type ReservationDao struct {
	gorm.Model
	applicantID uint
}

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&ReservationDao{})
	} else {
		fmt.Println("Error happened when initializing ResservationDao and creating Table in function NoticeDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	connectDatabase()
}

func InsertReservation(applicantID uint) error {
	var reservation ReservationDao
	reservation.applicantID = applicantID
	err := db.Transaction(func(tx *gorm.DB) error {
		if DBErr == nil {
			DBErr = tx.Create(&reservation).Error
		}
		return DBErr
	})
	if err != nil {
		fmt.Println("Error happened when inserting reservation in function ReservationDao.InsertReservation()")
		fmt.Println(err)
	}
	return err
}
