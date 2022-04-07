package NoticeService

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Dao/NoticeDao"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/Notice"
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
