package BillDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/Bill"
	"fmt"
	"gorm.io/gorm"
)

type BillDao struct {
	gorm.Model
	ReservationID uint    `gorm:"type:uint"`
	PayerID       uint    `gorm:"type:uint"`
	Money         float64 `gorm:"type:int"`
	BillStatus    int     `gorm:"type:int"`
}

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&BillDao{})
	} else {
		fmt.Println("Error happened when initializing MentorRecordDao and creating Table in function MentorRecordDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	connectDatabase()
}
func (BillDao) TableName() string {
	return "bills"
}
func convertBillToDao(bill Bill.Bill) BillDao {
	return BillDao{
		ReservationID: bill.ReservationID,
		PayerID:       bill.PayerID,
		Money:         bill.Money,
		BillStatus:    bill.BillStatus,
	}
}
func convertDaoToBill(dao BillDao) Bill.Bill {
	return Bill.Bill{
		BillID:        dao.ID,
		ReservationID: dao.ReservationID,
		PayerID:       dao.PayerID,
		Money:         dao.Money,
		BillStatus:    dao.BillStatus,
	}
}
func InsertBill(bill Bill.Bill) error {
	var billDao = convertBillToDao(bill)
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&billDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when inserting bills in function BillDao.InsertBill()")
	}
	return err
}
func FindBillsByPayerID(id uint) ([]Bill.Bill, error) {
	var bill []Bill.Bill
	var billDao []BillDao
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("payer_id=?", id).Find(&billDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding bills in function BillDao.FindBillByPayerID()")
	} else {
		bill = make([]Bill.Bill, len(billDao), len(billDao))
		for key := range billDao {
			bill[key] = convertDaoToBill(billDao[key])
		}
	}
	return bill, err
}
func FindBillByBillID(id uint) (Bill.Bill, error) {
	var bill Bill.Bill
	var billDao BillDao
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("id=?", id).Find(&billDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding bills in function BillDao.FindBillByPayerID()")
	} else {
		bill = convertDaoToBill(billDao)
	}
	return bill, err
}
func UpdateBillStatus(id uint, status int) error {
	var billDao BillDao
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Model(&billDao).Where("id=?", id).Update("bill_status", status).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when updating bills in function BillDao.UpdateBillStatus()")
	}
	return err
}
func FindBillByReservationID(id uint) (Bill.Bill, error) {
	var bill Bill.Bill
	var billDao BillDao
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where(&BillDao{ReservationID: id}).Find(&billDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding bills in function BillDao.FindBillByReservationID()")
	} else {
		bill = convertDaoToBill(billDao)
	}
	return bill, err
}
