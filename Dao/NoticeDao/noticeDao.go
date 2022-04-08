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
func FindNoticeByIssuerID(issuerid uint) ([]Notice.Notice, int, error) {
	var noticeDaos []NoticeDao
	var notices []Notice.Notice
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("issuer_id=?", issuerid).Find(&noticeDaos).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding notices in function NoticeDao.FindNoticeByIssuerID()")
	} else {
		for key := range noticeDaos {
			notices[key] = convertDaoToNotice(noticeDaos[key])
		}
	}
	return notices, len(notices), err
}
func FindNoticeByOffset(offset int, limit int) ([]Notice.Notice, int, error) {
	var count int64
	var daos = make([]NoticeDao, limit)
	var notices = make([]Notice.Notice, limit)
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Limit(limit).Offset(offset).Order("id").Find(&daos).Count(&count).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding users in function NoticeDao.FindAllNotice()")
		return notices, int(count), err
	}
	for key := range daos {
		notices[key] = convertDaoToNotice(daos[key])
	}
	return notices, int(count), nil
}
