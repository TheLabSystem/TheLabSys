package NoticeController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Service/NoticeService"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/NoticeService/NoticeRequestAndResponse"
	"net/http"

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
	response.Code, response.Data.Message, response.Data.Notice = NoticeService.NoticeService{}.GetNoticeList()
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
