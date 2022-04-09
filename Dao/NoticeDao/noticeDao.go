package NoticeDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/Notice"
	"fmt"

	"gorm.io/gorm"
)

type NoticeDao struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255)"`
	Content  string `gorm:"type:text(65535)"`
	IssuerID uint   `gorm:"type:integer"`
}

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&NoticeDao{})
	} else {
		fmt.Println("Error happened when initializing NoticeDao and creating Table in function NoticeDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	connectDatabase()
}
func (NoticeDao) TableName() string {
	return "notices"
}
func convertDaoToNotice(dao NoticeDao) Notice.Notice {
	return Notice.Notice{
		NoticeID: dao.ID,
		Title:    dao.Title,
		Content:  dao.Content,
		IssuerID: dao.IssuerID,
	}
}
func convertNoticeToDao(notice Notice.Notice) NoticeDao {
	return NoticeDao{
		Title:    notice.Title,
		Content:  notice.Content,
		IssuerID: notice.IssuerID,
	}
}
func InsertNotice(notice Notice.Notice) error {
	var noticeDao = convertNoticeToDao(notice)
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&noticeDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when inserting notices in function NoticeDao.InsertNotice()")
	}
	return err
}
func DeleteNotice(notice Notice.Notice) error {
	var noticeDao NoticeDao
	noticeDao.ID = notice.NoticeID
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Delete(&noticeDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when deleting notices in function NoticeDao.DeleteNotice()")
	}
	return err
}
func UpdateNotice(notice Notice.Notice) error {
	var noticeDao = convertNoticeToDao(notice)
	noticeDao.ID = notice.NoticeID
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Model(&noticeDao).Updates(noticeDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when updating notices in function NoticeDao.UpdateNotice()")
	}
	return err
}
func FindNoticeByID(id uint) (Notice.Notice, error) {
	var noticeDao NoticeDao
	var notice Notice.Notice
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("id=?", id).First(&noticeDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding notices in function NoticeDao.FindNoticeByID()")
	} else {
		notice = convertDaoToNotice(noticeDao)
	}
	return notice, err
}
func FindNoticeByOffset() ([]Notice.Notice, error) {
	var count int64
	var daos = make([]NoticeDao, 0)
	err1 := db.Transaction(
		func(tx *gorm.DB) error {
			if err1 := tx.Find(&daos).Count(&count).Error; err1 != nil {
				tx.Rollback()
				return err1
			}
			return nil
		})
	if err1 != nil {
		fmt.Println("Error happened when finding users in function NoticeDao.FindAllNotice()1")
		return make([]Notice.Notice, 0), err1
	}
	daos = make([]NoticeDao, count)
	var notices = make([]Notice.Notice, count)
	err2 := db.Transaction(
		func(tx *gorm.DB) error {
			if err2 := tx.Order("id").Find(&daos).Error; err2 != nil {
				tx.Rollback()
				return err2
			}
			return nil
		})
	if err2 != nil {
		fmt.Println("Error happened when finding users in function NoticeDao.FindAllNotice()2")
		return notices, err2
	}
	for key := 0; key < len(daos); key++ {
		fmt.Println(key, daos[key])
		notices[key] = convertDaoToNotice(daos[key])
	}
	return notices, nil
}
