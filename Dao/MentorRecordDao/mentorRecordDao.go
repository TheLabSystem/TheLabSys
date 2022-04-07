package MentorRecordDao

import (
	"TheLabSystem/Dao/DBAccessor"
	"TheLabSystem/Types/ServiceType/MentorRecord"
	"fmt"
	"gorm.io/gorm"
)

type MentorRecordDao struct {
	gorm.Model
	StudentID uint `gorm:"type:integer"`
	TeacherID uint `gorm:"type:integer"`
}

var db *gorm.DB
var DBErr error

func connectDatabase() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		DBErr = db.AutoMigrate(&MentorRecordDao{})
	} else {
		fmt.Println("Error happened when initializing MentorRecordDao and creating Table in function MentorRecordDao.connectDatabase()")
		fmt.Println(DBErr)
	}
}
func init() {
	connectDatabase()
}
func (MentorRecordDao) TableName() string {
	return "mentorRecord"
}
func convertDaoToMentorRecord(dao MentorRecordDao) MentorRecord.MentorRecord {
	return MentorRecord.MentorRecord{
		StudentID: dao.StudentID,
		TeacherID: dao.TeacherID,
	}
}
func convertMentorRecordToDao(mr MentorRecord.MentorRecord) MentorRecordDao {
	return MentorRecordDao{
		StudentID: mr.StudentID,
		TeacherID: mr.TeacherID,
	}
}
func InsertMentorRecord(mr MentorRecord.MentorRecord) error {
	var mrDao = convertMentorRecordToDao(mr)
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&mrDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when inserting mentorRecords in function MentorRecordDao.InsertMentorRecord()")
		fmt.Println(err)
	}
	return err
}
func DeleteMentorRecord(mr MentorRecord.MentorRecord) error {
	var mrDao = convertMentorRecordToDao(mr)
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("student_id=? and teacher_id=?", mr.StudentID, mr.TeacherID).Delete(&mrDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when deleting mentorRecords in function MentorRecordDao.DeleteMentorRecord()")
	}
	return err
}
func FindMentorRecordByStudentID(id uint) (MentorRecord.MentorRecord, error) {
	var mrDao MentorRecordDao
	var mr MentorRecord.MentorRecord
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("student_id=?", id).First(&mrDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding mentorRecords in function MentorRecordDao.FindMentorRecordByStudentID()")
	} else {
		mr = convertDaoToMentorRecord(mrDao)
	}
	return mr, err
}
func FindMentorRecordByTeacherID(id uint) (MentorRecord.MentorRecord, error) {
	var mrDao MentorRecordDao
	var mr MentorRecord.MentorRecord
	err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("teacher_id=?", id).First(&mrDao).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	if err != nil {
		fmt.Println("Error happened when finding mentorRecords in function MentorRecordDao.FindMentorRecordByTeacherID()")
	} else {
		mr = convertDaoToMentorRecord(mrDao)
	}
	return mr, err
}
