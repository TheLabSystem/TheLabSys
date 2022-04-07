package AuthController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Service/UserService"
	"TheLabSystem/Types/RequestAndResponseType/Auth/LoginRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Auth/LogoutRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Auth/WhoAmIRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
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
	loginResponse.Data.Message = ErrorInformation.GenerateErrorInformation(loginResponse.Code)
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
