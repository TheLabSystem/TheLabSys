package AuthController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Service/UserService"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/LoginRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/LogoutRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/WhoAmIRequestAndResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
}

func (controller AuthController) Login(c *gin.Context) {
	loginRequest := &LoginRequestAndResponse.LoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// checkUserPassword
	var userService UserService.UserService
	errNo, user, userinfo := userService.CheckPassword(loginRequest.Username, loginRequest.Password)
	loginResponse := &LoginRequestAndResponse.LoginResponse{}
	loginResponse.Code = errNo
	loginResponse.Data.Userinfo = userinfo
	loginResponse.Data.User = user
	if errNo == ErrNo.OK {
		loginResponse.Data.Message = "Log in successfully"
	} else if errNo == ErrNo.WrongPassword {
		loginResponse.Data.Message = "Wrong Password!"
	} else if errNo == ErrNo.UnknownError {
		loginResponse.Data.Message = "An unexpected error has happened.Please try again."
	}
	c.SetCookie("camp-session", loginRequest.Username, 0, "/", "/", false, false)
	c.JSON(http.StatusOK, loginResponse)
}

func (controller AuthController) Logout(c *gin.Context) {
	logoutResponse := &LogoutRequestAndResponse.LogoutResponse{}

	cookie, err := c.Request.Cookie("camp-session")
	if err == nil {
		c.SetCookie(cookie.Name, cookie.Value, -1, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		logoutResponse.Code = ErrNo.OK
	} else {
		logoutResponse.Code = ErrNo.LoginRequired
	}
	c.JSON(http.StatusOK, logoutResponse)
}

func (controller AuthController) WhoAmI(c *gin.Context) {
	whoAmIResponse := WhoAmIRequestAndResponse.WhoAmIResponse{}
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		whoAmIResponse.Code = ErrNo.LoginRequired
		whoAmIResponse.Data.Message = ErrorInformation.GenerateErrorInformation(whoAmIResponse.Code)
		c.JSON(http.StatusOK, whoAmIResponse)
		return
	}
	whoAmIResponse.Data.User, err = UserDao.FindUserByUsername(cookie.Value)
	if err != nil {
		whoAmIResponse.Code = ErrNo.UnknownError
		whoAmIResponse.Data.Message = ErrorInformation.GenerateErrorInformation(whoAmIResponse.Code)
	} else if whoAmIResponse.Data.User.UserID == 0 {
		whoAmIResponse.Code = ErrNo.UserNotExisted
		whoAmIResponse.Data.Message = ErrorInformation.GenerateErrorInformation(whoAmIResponse.Code)
	} else {
		whoAmIResponse.Code = ErrNo.OK
		whoAmIResponse.Data.Message = ErrorInformation.GenerateErrorInformation(whoAmIResponse.Code)
	}
	c.JSON(http.StatusOK, whoAmIResponse)
}
