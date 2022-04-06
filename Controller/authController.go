package Controller

import (
	"TheLabSystem/Service/UserService"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/LoginRequestAndResponse"
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
	c.JSON(http.StatusOK, loginResponse)
}
