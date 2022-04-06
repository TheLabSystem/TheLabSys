package UserDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/User"
	"fmt"
	"gorm.io/gorm"
)

type UserDao struct {
	gorm.Model
	Username    string `gorm:"type:varchar(25)"`
	UserType    int    `gorm:"type:integer"`
	DisplayName string `gorm:"type:varchar(25)"`
	Password    string `gorm:"type:varchar(30)"`
}

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&UserDao{})
	} else {
		fmt.Println("Error happened when initializing UserDao and creating Table in function UserDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	// if table does not exist, then create table
	connectDatabase()
	//for DBErr != nil {
	//	time.Sleep(30000)
	//	connectDatabase()
	//}
}
func (UserDao) TableName() string {
	return "users"
}
func convertDaoToUser(dao UserDao) User.User {
	return User.User{
		UserID:      dao.ID,
		Username:    dao.Username,
		UserType:    dao.UserType,
		DisplayName: dao.DisplayName,
		Password:    dao.Password,
	}
}
func convertUserToDao(user User.User) UserDao {
	return UserDao{
		UserType:    user.UserType,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Password:    user.Password,
	}
}
func InsertUser(user User.User) error {
	var userDao = convertUserToDao(user)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&userDao).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error happened when inserting users in function UserDao.InsertUser()")
	}
	return err
}
func DeleteUser(user User.User) error {
	var userDao UserDao
	userDao.ID = user.UserID
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&userDao).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error happened when deleting users in function UserDao.DeleteUser()")
	}
	return err
}

// UpdateUser update User by ID
func UpdateUser(user User.User) error {
	var userDao = convertUserToDao(user)
	userDao.ID = user.UserID
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&userDao).Updates(userDao).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error happened when updating users in function UserDao.UpdateUserByID()")
	}
	return err
}

func FindUserByID(id uint) (User.User, error) {
	var userDao UserDao
	var user User.User
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).First(&userDao).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error happened when finding users in function UserDao.FindUserByID()")
	} else {
		user = convertDaoToUser(userDao)
	}
	return user, err
}
func FindUserByUsername(username string) (User.User, error) {
	var userDao UserDao
	var user User.User
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("username = ?", username).First(&userDao).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error happened when finding users in function UserDao.FindUserByUsername()")
	} else {
		user = convertDaoToUser(userDao)
	}
	return user, err
}
func FindAllUser(offset int, limit int) ([]User.User, error) {
	var daos = make([]UserDao, limit, limit)
	var users = make([]User.User, limit, limit)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Limit(limit).Offset(offset).Order("id").Find(&daos).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error happened when finding users in function UserDao.FindAllUser()")
		return users, err
	}
	for key := range daos {
		users[key] = convertDaoToUser(daos[key])
	}
	return users, nil
}