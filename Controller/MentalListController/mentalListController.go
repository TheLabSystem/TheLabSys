package MentalListController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Service/MentalListService"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/MentalList/AddStudentRequestAndResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MentalListController struct {
}

func (controller MentalListController) AddStudentController(c *gin.Context) {
	var request = &AddStudentRequestAndResponse.AddStudentRequest{}
	var response = &AddStudentRequestAndResponse.AddStudentResponse{}
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
	response.Code = MentalListService.MentalListService{}.AddStudent(request.StudentID, cookie)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
}
