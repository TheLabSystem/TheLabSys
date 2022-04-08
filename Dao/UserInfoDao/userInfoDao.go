package UserInfoDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserInfoDao struct {
	gorm.Model
	UserInfo string `gorm:"type:text(65535)"`
	UserID   uint   `gorm:"type:integer"`
}

var db *gorm.DB
var DBErr error

func (UserInfoDao) TableName() string {
	return "userinfo"
}
func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&UserInfoDao{})
	} else {
		fmt.Println("Error happened when initializing UserInfoDao and creating Table in function UserInfoDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	// if table does not exist, then create table
	connectDatabase()
	//for DBErr != nil {
	//	time.Sleep(300000)
	//	connectDatabase()
	//}
}
func ChangeUserInfo(userID uint, info string) error {
	var userInfoDao UserInfoDao
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id=?", userID).First(&userInfoDao).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	userInfoDao.UserID = userID
	userInfoDao.UserInfo = info
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&userInfoDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	} else {
		err = db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&userInfoDao).Updates(&userInfoDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
		fmt.Println()
	}
	if err != nil {
		fmt.Println("Error happened when changing userinfo in function UserInfoDao.ChangeUserInfo()")
		fmt.Println(err)
	}
	return err
}
func FindUserInfoByID(userID uint) (string, error) {
	var userInfoDao UserInfoDao
	userInfoDao.UserID = userID
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id=?", userID).First(&userInfoDao).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil
	}
	if err != nil {
		fmt.Println("Error happened when finding userinfo in function UserInfoDao.FindUserByID()")
		fmt.Println(err)
		return "", err
	}
	return userInfoDao.UserInfo, err
}
