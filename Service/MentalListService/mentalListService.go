package MentalListService

import (
	"TheLabSystem/Dao/MentorRecordDao"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/MentorRecord"
	"TheLabSystem/Types/ServiceType/User"
)

type MentalListService struct {
}

func (service MentalListService) AddStudent(studentID uint, username string) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username == "" {
		return ErrNo.LoginRequired
	}
	if user.UserType != 3 {
		return ErrNo.PermDenied
	}
	var student User.User
	if student, err = UserDao.FindUserByID(studentID); err != nil {
		return ErrNo.UnknownError
	} else if student.Username == "" {
		return ErrNo.StudentNotExist
	}
	mentor := MentorRecord.MentorRecord{
		StudentID: studentID,
		TeacherID: user.UserID,
	}
	if MentorRecordDao.InsertMentorRecord(mentor) != nil {
		return ErrNo.UnknownError
	} else {
		return ErrNo.OK
	}
}

func (service MentalListService) DeleteStudent(studentID uint, username string) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username == "" {
		return ErrNo.LoginRequired
	}
	if user.UserType != 3 {
		return ErrNo.PermDenied
	}
	mentor := MentorRecord.MentorRecord{
		StudentID: studentID,
		TeacherID: user.UserID,
	}
	if MentorRecordDao.InsertMentorRecord(mentor) != nil {
		return ErrNo.UnknownError
	} else {
		return ErrNo.OK
	}
}

func (service MentalListService) ViewStudent(username string) ([]MentorRecord.MentorRecord, ErrNo.ErrNo) {
	var res []MentorRecord.MentorRecord
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return res, ErrNo.UnknownError
	}
	if user.Username == "" {
		return res, ErrNo.LoginRequired
	}
	if user.UserType != 3 {
		return res, ErrNo.PermDenied
	}
	res, err = MentorRecordDao.FindMentorRecordByTeacherID(user.UserID)
	if err != nil {
		return res, ErrNo.UnknownError
	} else {
		return res, ErrNo.OK
	}
}
