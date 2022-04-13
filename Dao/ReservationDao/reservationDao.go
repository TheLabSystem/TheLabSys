package ReservationDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Dao/ReservationInfoDao"
	"TheLabSystem/Types/ServiceType/MentorRecord"
	"TheLabSystem/Types/ServiceType/Reservation"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type ReservationDao struct {
	gorm.Model
	ApplicantID uint `gorm:"type:uint"`
	Status      int  `gorm:"type:int"`
}

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&ReservationDao{})
	} else {
		fmt.Println("Error happened when initializing ReservationDao and creating Table in function NoticeDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	connectDatabase()
}
func (ReservationDao) TableName() string {
	return "reservation"
}
func convertReservationToDao(reservation Reservation.Reservation) ReservationDao {
	return ReservationDao{
		ApplicantID: reservation.ApplicantID,
		Status:      reservation.Status,
	}
}
func convertDaoToReservation(reservationDao ReservationDao) Reservation.Reservation {
	return Reservation.Reservation{
		ReservationID: reservationDao.ID,
		ApplicantID:   reservationDao.ApplicantID,
		OperatingDay:  reservationDao.CreatedAt.Format("2006-01-02"),
		Status:        reservationDao.Status,
	}
}
func InsertReservation(reservation Reservation.Reservation) (Reservation.Reservation, error) {
	var reservationDao = convertReservationToDao(reservation)
	var res Reservation.Reservation
	err := db.Transaction(func(tx *gorm.DB) error {
		if DBErr == nil {
			DBErr = tx.Create(&reservationDao).Error
		}
		return DBErr
	})
	if err != nil {
		fmt.Println("Error happened when inserting reservation in function ReservationDao.InsertReservation()")
		fmt.Println(err)
		return res, err
	}
	res = convertDaoToReservation(reservationDao)
	return res, err
}

func UpdateReservation(reservationID uint, status int) error {
	var reservationDao ReservationDao
	reservationDao.ID = reservationID
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Model(&reservationDao).Update("status", status).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("更新预约状态出错")
	}
	return err
}

func FindReservationByID(id uint) (Reservation.Reservation, error) {
	var reservationDao ReservationDao
	var reservation Reservation.Reservation
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("id=?", id).First(&reservationDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("查找预约出现错误!")
	} else {
		reservation = convertDaoToReservation(reservationDao)
	}
	return reservation, err
}

func FindAllReservation() ([]Reservation.Reservation, error) {
	var reservationDao []ReservationDao
	var reservation []Reservation.Reservation
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Find(&reservationDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("查找预约出现错误")
	} else {
		reservation = make([]Reservation.Reservation, len(reservationDao), len(reservationDao))
		for key := range reservationDao {
			reservation[key] = convertDaoToReservation(reservationDao[key])
		}
	}
	return reservation, err
}
func FindApprovalReservation(userType int, mentorRecord []MentorRecord.MentorRecord) ([]Reservation.Reservation, error) {
	var reservationDao []ReservationDao
	var reservation []Reservation.Reservation
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("status!=1 and status!=2 and status!=3").Find(&reservationDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("查找预约出现错误")
	} else {
		var i = 0
		var flag bool
		for key := 0; key < len(reservationDao); key++ {
			info, err := ReservationInfoDao.FindInfoByReservationID(reservationDao[key].ID)
			if err != nil {
				fmt.Println("error!")
			}
			flag = true
			for j := 0; j < len(info); j++ {
				day, _ := time.Parse("2006-01-02", info[j].ReservationDay)
				if day.Before(time.Now()) {
					flag = false
				}
			}
			if flag == true {
				if userType == 3 {
					if reservationDao[key].Status == 112 {
						var isMentor bool
						isMentor = false
						for b := 0; b < len(mentorRecord); b++ {
							if reservationDao[key].ApplicantID == mentorRecord[b].StudentID {
								isMentor = true
								break
							}
						}
						if isMentor {
							reservation = append(reservation, convertDaoToReservation(reservationDao[key]))
							i++
						}
					}
				} else if userType == 4 {
					if reservationDao[key].Status == 12 || reservationDao[key].Status == 21234 || reservationDao[key].Status == 24 || reservationDao[key].Status == 32 {
						reservation = append(reservation, convertDaoToReservation(reservationDao[key]))
						i++
					}
				} else if userType == 255 {
					if reservationDao[key].Status == 234 {
						reservation = append(reservation, convertDaoToReservation(reservationDao[key]))
						i++
					}
				}
			}
		}
	}
	return reservation, err
}

func FindReservationByApplicantID(id uint) ([]Reservation.Reservation, error) {
	var reservationDaos []ReservationDao
	var reservations []Reservation.Reservation
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where(&ReservationDao{ApplicantID: id}).Find(&reservationDaos).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return reservations, nil
	} else if err != nil {
		fmt.Println("Error happened when finding reservation by applicantID in function FindReservationByApplicantID")
	} else {
		reservations = make([]Reservation.Reservation, len(reservationDaos), len(reservationDaos))
		for key := range reservationDaos {
			reservations[key] = convertDaoToReservation(reservationDaos[key])
		}
	}
	return reservations, err
}
