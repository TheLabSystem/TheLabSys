package UserService

import (
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Dao/UserInfoDao"
	"TheLabSystem/Dao/VerifyCodeDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/User"
	"fmt"
)

type UserService struct {
}

func (userService UserService) CheckPassword(username string, password string) (ErrNo.ErrNo, User.User, string) {
	user, err := UserDao.FindUserByUsername(username)
	fmt.Println(user)
	if err != nil {
		fmt.Println("Error happened when check user's password in function checkPassword.")
		return ErrNo.UnknownError, User.User{}, ""
	} else if user.Password != password {
		return ErrNo.WrongPassword, User.User{}, ""
	} else {
		info, _ := UserInfoDao.FindUserInfoByID(user.UserID)
		return ErrNo.OK, user, info
	}
}
func (userService UserService) ChangeUserInfo(username string, password string, displayName string, userInfo string) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if displayName != "" {
		user.DisplayName = displayName
	}
	if password != "" {
		user.Password = password
	}
	if UserDao.UpdateUser(user) != nil {
		return ErrNo.UnknownError
	}
	if userInfo != "" {
		if UserInfoDao.ChangeUserInfo(user.UserID, userInfo) != nil {
			return ErrNo.UnknownError
		}
	}
	return ErrNo.OK
}
func (userService UserService) RegisterUser(username string, password string, userType int, verifyCode int) ErrNo.ErrNo {
	verify, err := VerifyCodeDao.CheckVerifyCode(verifyCode, userType)
	if err != nil {
		return ErrNo.UnknownError
	} else if verify == false {
		return ErrNo.VerifyCodeNotValid
	}
	var user User.User
	user, err = UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username != "" {
		return ErrNo.UserHasExisted
	}
	user.Username = username
	user.DisplayName = username
	user.UserType = userType
	user.Money = 0
	user.Password = password
	if err = UserDao.InsertUser(user); err != nil {
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
func (userService UserService) FindUserInfo(username string) (ErrNo.ErrNo, User.User, string) {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError, user, ""
	}
	if user.Username == "" {
		return ErrNo.UserNotExisted, user, ""
	}
	var info string
	info, err = UserInfoDao.FindUserInfoByID(user.UserID)
	if err != nil {
		return ErrNo.UnknownError, user, info
	}
	return ErrNo.OK, user, info
}
func (userService UserService) AddMoneyService(money float64, username string) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username == "" {
		return ErrNo.UserNotExisted
	}
	user.Money = user.Money + money
	err = UserDao.UpdateUser(user)
	if err != nil {
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
