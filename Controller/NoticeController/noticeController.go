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
