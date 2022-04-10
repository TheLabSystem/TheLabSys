package ReportFormController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Service/ReportFormService"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/ReportForm/GetReportFormRequestAndResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReportFormController struct {
}

func (controller ReportFormController) GetReportForm(c *gin.Context) {
	request := &GetReportFormRequestAndResponse.GetReportFormRequest{}
	response := &GetReportFormRequestAndResponse.GetReportFormResponse{}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	if err = c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response.Data.ReportForm, response.Code = ReportFormService.ReportFormService{}.GetReportForm(request.StartDate, request.EndDate, cookie)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
}
