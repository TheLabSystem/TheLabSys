package BillController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Service/BillService"
	"TheLabSystem/Types/RequestAndResponseType/Bill/GetBillRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BillController struct {
}

// 需要完成payBill，getBill的

func (controller BillController) GetBill(c *gin.Context) {
	var request = &GetBillRequestAndResponse.GetBillRequest{}
	var response = &GetBillRequestAndResponse.GetBillResponse{}
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
	response.Data.Bills, response.Code = BillService.BillService{}.GetBill(cookie)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
	return
}
