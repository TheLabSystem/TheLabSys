package NoticeService

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Dao/NoticeDao"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/Notice"
	"fmt"
)

type NoticeService struct {
}

func (noticeService NoticeService) AddNotice(cookie string, title string, content string) (ErrNo.ErrNo, string) {
	user, err := UserDao.FindUserByUsername(cookie)
	if err != nil {
		return ErrNo.UnknownError, ErrorInformation.GenerateErrorInformation(ErrNo.UnknownError)
	}
	if title == "" {
		return ErrNo.ParamInvalid, "标题不能为空"
	}
	if content == "" {
		return ErrNo.ParamInvalid, "内容不能为空"
	}
	if NoticeDao.InsertNotice(Notice.Notice{
		Title:    title,
		Content:  content,
		IssuerID: user.UserID,
	}) != nil {
		return ErrNo.UnknownError, ErrorInformation.GenerateErrorInformation(ErrNo.UnknownError)
	}
	return ErrNo.OK, ErrorInformation.GenerateErrorInformation(ErrNo.OK)
}

func (noticeService NoticeService) GetNoticeList(offset int, limit int) (ErrNo.ErrNo, string, []Notice.Notice, int) {
	noticeList, total, err := NoticeDao.FindNoticeByOffset(offset, limit)
	if err != nil {
		fmt.Println(err)
		return ErrNo.UnknownError, ErrorInformation.GenerateErrorInformation(ErrNo.UnknownError), noticeList, total
	}
	return ErrNo.OK, ErrorInformation.GenerateErrorInformation(ErrNo.OK), noticeList, total
}

func (noticeService NoticeService) GetNoticeListByIssuer(issuerId int) (ErrNo.ErrNo, string, []Notice.Notice, int) {
	noticeList, total, err := NoticeDao.FindNoticeByIssuerID(uint(issuerId))
	if err != nil {
		return ErrNo.UnknownError, ErrorInformation.GenerateErrorInformation(ErrNo.UnknownError), noticeList, 0
	}
	return ErrNo.OK, ErrorInformation.GenerateErrorInformation(ErrNo.OK), noticeList, total
}

func (noticeService NoticeService) DeleteNotice(username string, noticeId int) (ErrNo.ErrNo, string) {
	user, userErr := UserDao.FindUserByUsername(username)
	notice, noticeErr := NoticeDao.FindNoticeByID(uint(noticeId))
	if userErr != nil || noticeErr != nil {
		return ErrNo.UnknownError, ErrorInformation.GenerateErrorInformation(ErrNo.UnknownError)
	}
	if user.UserID != notice.IssuerID {
		return ErrNo.PermDenied, ErrorInformation.GenerateErrorInformation(ErrNo.PermDenied)
	}
	if NoticeDao.DeleteNotice(notice) != nil {
		return ErrNo.UnknownError, ErrorInformation.GenerateErrorInformation(ErrNo.UnknownError)
	}
	return ErrNo.OK, ErrorInformation.GenerateErrorInformation(ErrNo.OK)
}
