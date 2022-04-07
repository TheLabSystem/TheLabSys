package VerifyCodeService

import (
	"TheLabSystem/Config/UserPermissionDecide"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Dao/VerifyCodeDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/VerifyCode"
)

type VerifyCodeService struct {
}

func (service VerifyCodeService) AddVerifyCode(code int, usertype int, username string) ErrNo.ErrNo {
	// check permission
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

func (service VerifyCodeService) ViewAllVerifyCode(username string) ([]VerifyCode.VerifyCode, ErrNo.ErrNo) {
	// check permission
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return nil, ErrNo.UnknownError
	}
	if user.Username == "" {
		return nil, ErrNo.LoginRequired
	}
	if UserPermissionDecide.ViewVerifyCode(user.UserType) {
		res, err := VerifyCodeDao.ViewVerifyCode()
		if err != nil {
			return nil, ErrNo.UnknownError
		}
		return res, ErrNo.OK
	} else {
		return nil, ErrNo.PermDenied
	}
}

func (service VerifyCodeService) DeleteVerifyCode(code int, username string) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if UserPermissionDecide.DeleteVerifyCode(user.UserType) {
		err := VerifyCodeDao.DeleteVerifyCode(code)
		if err != nil {
			return ErrNo.UnknownError
		}
		return ErrNo.OK
	} else {
		return ErrNo.PermDenied
	}
}
