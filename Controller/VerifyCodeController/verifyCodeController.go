package VerifyCodeController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Service/VerifyCodeService"
	"TheLabSystem/Types/RequestAndResponseType/AddVerifyCodeRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VerifyCodeController struct {
}

func (controller VerifyCodeController) AddVerifyCodeController(c *gin.Context) {
	var request = &AddVerifyCodeRequestAndResponse.AddVerifyCodeRequest{}
	var response = &AddVerifyCodeRequestAndResponse.AddVerifyCodeResponse{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = VerifyCodeService.VerifyCodeService{}.AddVerifyCode(request.VerifyCode, request.UserType, cookie)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
	return
}
