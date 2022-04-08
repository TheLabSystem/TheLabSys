package UserServiceController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Service/UserService"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/UserService/AddMoneyServiceRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/UserService/ChangeUserInfoRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/UserService/FindUserInfoRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/UserService/RegisterUserRequestAndResponse"
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

func (controller UserServiceController) FindUserInfo(c *gin.Context) {
	request := &FindUserInfoRequestAndResponse.FindUserInfoRequest{}
	response := &FindUserInfoRequestAndResponse.FindUserInfoResponse{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code, response.Data.User, response.Data.UserInfo = UserService.UserService{}.FindUserInfo(cookie)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
}

func (controller UserServiceController) AddMoney(c *gin.Context) {
	request := &AddMoneyServiceRequestAndResponse.AddMoneyServiceRequest{}
	response := &AddMoneyServiceRequestAndResponse.AddMoneyServiceResponse{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = UserService.UserService{}.AddMoneyService(request.Money, cookie)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
}
