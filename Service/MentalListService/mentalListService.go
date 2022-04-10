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
	if MentorRecordDao.DeleteMentorRecord(mentor) != nil {
		return ErrNo.UnknownError
	} else {
		return ErrNo.OK
	}
}

func (service MentalListService) ViewStudent(username string) ([]User.User, ErrNo.ErrNo) {
	var errA error
	var res []User.User
	var rec []MentorRecord.MentorRecord
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return res, ErrNo.UnknownError
	}
	rec, errB := MentorRecordDao.FindMentorRecordByTeacherID(user.UserID)
	if errB != nil {
		return res, ErrNo.UnknownError
	}
	res = make([]User.User, len(rec), len(rec))
	for key := range rec {
		res[key], errA = UserDao.FindUserByID(rec[key].StudentID)
		if errA != nil {
			return res, ErrNo.UnknownError
		}
	}
	if user.Username == "" {
		return res, ErrNo.LoginRequired
	}
	if user.UserType != 3 {
		return res, ErrNo.PermDenied
	}
	if err != nil {
		return res, ErrNo.UnknownError
	} else {
		return res, ErrNo.OK
	}
}
