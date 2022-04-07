package UserServiceController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Service/UserService"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/UerService/ChangeUserInfoRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/UerService/RegisterUserRequestAndResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserServiceController struct {
}

func (controller UserServiceController) ChangeUserInfo(c *gin.Context) {
	response := &ChangeUserInfoRequestAndResponse.ChangeUserInfoResponse{}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	request := &ChangeUserInfoRequestAndResponse.ChangeUserInfoRequest{}
	if err = c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response.Code = UserService.UserService{}.ChangeUserInfo(cookie, request.NewPassword, request.DisplayName, request.UserInfo)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
}

func (controller UserServiceController) RegisterUser(c *gin.Context) {
	request := &RegisterUserRequestAndResponse.RegisterUserRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response := &RegisterUserRequestAndResponse.RegisterUserResponse{}
	response.Code = UserService.UserService{}.RegisterUser(request.Username, request.Password, request.UserType, request.VerifyCode)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
}
