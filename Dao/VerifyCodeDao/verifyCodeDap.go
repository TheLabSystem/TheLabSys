package VerifyCodeDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/VerifyCode"
	"fmt"
	"gorm.io/gorm"
)

type VerifyCodeDao struct {
	gorm.Model
	verifyCode int `gorm:"type(integer)"`
	usertype   int `gorm:"type(integer)"`
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
func InsertVerifyCode(code int, userType int) error {
	var verifyCodeDao = VerifyCodeDao{
		verifyCode: code,
		usertype:   userType,
	}
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&verifyCodeDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when inserting verifyCode in function VerifyCodeDao.InsertVerifyCode()")
	}
	return err
}
func DeleteVerifyCode(code int) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(&VerifyCodeDao{verifyCode: code}).Delete(VerifyCodeDao{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error happened when deleting verifyCode in function VerifyCode.DeleteVerifyCode()")
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
		fmt.Println("Error happened when viewing verifyCode in function VerifyCode.ViewVerifyCode()")
		fmt.Println(err)
	}
	sliceLen := len(daos)
	res = make([]VerifyCode.VerifyCode, sliceLen, sliceLen)
	return res, err
}
func CheckVerifyCode(code int, userType int) (bool, error) {
	var daos []VerifyCodeDao
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(&VerifyCodeDao{verifyCode: code}).Find(&daos).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error happened when checking verifyCode in function VerifyCode.CheckVerifyCode()")
		fmt.Println(err)
		return false, err
	}
	for key := range daos {
		if daos[key].usertype == userType {
			return true, err
		}
	}
	return false, err
}
