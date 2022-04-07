package VerifyCodeService

import (
	"TheLabSystem/Config/UserPermissionDecide"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Dao/VerifyCodeDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
)

type VerifyCodeService struct {
}

func (service VerifyCodeService) AddVerifyCode(code int, usertype int, username string) ErrNo.ErrNo {
	// check password
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username == "" {
		return ErrNo.LoginRequired
	}
	if UserPermissionDecide.AddVerifyCode(user.UserType) {
		err = VerifyCodeDao.InsertVerifyCode(code, usertype)
		if err != nil {
			return ErrNo.UnknownError
		}
		return ErrNo.OK
	} else {
		return ErrNo.PermDenied
	}
}
