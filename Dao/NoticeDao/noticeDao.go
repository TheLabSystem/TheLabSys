package NoticeDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/Notice"
	"fmt"
	"gorm.io/gorm"
)

type NoticeDao struct {
	gorm.Model
	NoticeInfo string `gorm:"type:text(65535)"`
	IssuerID   int    `gorm:"type:integer"`
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
		NoticeID:   dao.ID,
		NoticeInfo: dao.NoticeInfo,
		IssuerID:   dao.IssuerID,
	}
}
func convertNoticeToDao(notice Notice.Notice) NoticeDao {
	return NoticeDao{
		NoticeInfo: notice.NoticeInfo,
		IssuerID:   notice.IssuerID,
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
func FindNoticeByIssuerID(issuerid uint) (Notice.Notice, error) {
	var noticeDao NoticeDao
	var notice Notice.Notice
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("issuer_id=?", issuerid).First(&noticeDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding notices in function NoticeDao.FindNoticeByIssuerID()")
	} else {
		notice = convertDaoToNotice(noticeDao)
	}
	return notice, err
}
func FindNoticeByOffset(offset int, limit int) ([]Notice.Notice, error) {
	var daos = make([]NoticeDao, limit, limit)
	var notices = make([]Notice.Notice, limit, limit)
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Limit(limit).Offset(offset).Order("id").Find(&daos).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding users in function NoticeDao.FindAllNotice()")
		return notices, err
	}
	for key := range daos {
		notices[key] = convertDaoToNotice(daos[key])
	}
	return notices, nil
}
