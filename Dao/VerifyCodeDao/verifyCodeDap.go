package VerifyCodeDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/VerifyCode"
	"fmt"
	"gorm.io/gorm"
)

type VerifyCodeDao struct {
	gorm.Model
	VerifyCode int `gorm:"type(integer)"`
	UserType   int `gorm:"type(integer)"`
}

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&VerifyCodeDao{})
	} else {
		fmt.Println("Error happened when initializing VerifyCodeDao and creating Table in function VerifyCodeDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	connectDatabase()
}
func (VerifyCodeDao) TableName() string {
	return "verify_code"
}
func InsertVerifyCode(code int, userType int) error {
	var verifyCodeDao = VerifyCodeDao{
		VerifyCode: code,
		UserType:   userType,
	}
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&verifyCodeDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			fmt.Println(verifyCodeDao)
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when inserting VerifyCode in function VerifyCodeDao.InsertVerifyCode()")
	}
	return err
}
func DeleteVerifyCode(code int) error {
	fmt.Println(code)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("verify_code=?", code).Delete(&VerifyCodeDao{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error happened when deleting VerifyCode in function VerifyCode.DeleteVerifyCodeRequestAndResponse()")
		fmt.Println(err)
	}
	return nil
}
func ViewVerifyCode() ([]VerifyCode.VerifyCode, error) {
	var res []VerifyCode.VerifyCode
	var daos []VerifyCodeDao
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&daos).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error happened when viewing VerifyCode in function VerifyCode.ViewVerifyCode()")
		fmt.Println(err)
	}
	sliceLen := len(daos)
	res = make([]VerifyCode.VerifyCode, sliceLen, sliceLen)
	for key := range daos {
		res[key].VerifyCode = daos[key].VerifyCode
		res[key].UserType = daos[key].UserType
	}
	return res, err
}
func CheckVerifyCode(code int, userType int) (bool, error) {
	var daos []VerifyCodeDao
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(&VerifyCodeDao{VerifyCode: code}).Find(&daos).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error happened when checking VerifyCode in function VerifyCode.CheckVerifyCode()")
		fmt.Println(err)
		return false, err
	}
	for key := range daos {
		if daos[key].UserType == userType {
			return true, err
		}
	}
	return false, err
}
