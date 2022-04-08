package NoticeController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Service/NoticeService"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/NoticeService/NoticeRequestAndResponse"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoticeController struct {
}

func (controller NoticeController) AddNotice(c *gin.Context) {
	response := &NoticeRequestAndResponse.AddNoticeResponse{}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	request := &NoticeRequestAndResponse.AddNoticeRequest{}
	if err = c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response.Code, response.Data.Message = NoticeService.NoticeService{}.AddNotice(cookie, request.Title, request.Content)
	c.JSON(http.StatusOK, response)
}

func (controller NoticeController) GetNoticeList(c *gin.Context) {
	response := &NoticeRequestAndResponse.GetNoticeResponse{}
	issuerId, issuerIdErr := strconv.Atoi(c.DefaultQuery("issuerId", "-1"))
	offset, offsetErr := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, limitErr := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if issuerIdErr != nil || offsetErr != nil || limitErr != nil {
		response.Code = ErrNo.ParamInvalid
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	if issuerId == -1 {
		response.Code, response.Data.Message, response.Data.Notice, response.Data.Total = NoticeService.NoticeService{}.GetNoticeList(offset, limit)
	} else {
		response.Code, response.Data.Message, response.Data.Notice, response.Data.Total = NoticeService.NoticeService{}.GetNoticeListByIssuer(issuerId)
	}
	c.JSON(http.StatusOK, response)
}

func (controller NoticeController) DeleteNotice(c *gin.Context) {
	response := &NoticeRequestAndResponse.DeleteNoticeResponse{}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	request := &NoticeRequestAndResponse.DeleteNoticeRequest{}
	if err = c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response.Code, response.Data.Message = NoticeService.NoticeService{}.DeleteNotice(cookie, request.ID)
	c.JSON(http.StatusOK, response)
}

func (controller NoticeController) UpdateNotice(c *gin.Context) {
	response := &NoticeRequestAndResponse.UpdateNoticeResponse{}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	request := &NoticeRequestAndResponse.UpdateNoticeRequest{}
	if err = c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response.Code, response.Data.Message = NoticeService.NoticeService{}.UpdateNotice(cookie, request.ID, request.Title, request.Content)
	c.JSON(http.StatusOK, response)
}
