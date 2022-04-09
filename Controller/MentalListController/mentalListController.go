package MentalListController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Service/MentalListService"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/MentalList/AddStudentRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/MentalList/DeleteStudentRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/MentalList/ViewStudentRequestAndResponse"
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

func (controller MentalListController) DeleteStudentController(c *gin.Context) {
	var request = &DeleteStudentRequestAndResponse.DeleteStudentRequest{}
	var response = &DeleteStudentRequestAndResponse.DeleteStudentResponse{}
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
	response.Code = MentalListService.MentalListService{}.DeleteStudent(request.StudentID, cookie)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
}

func (controller MentalListController) ViewStudentController(c *gin.Context) {
	var request = &ViewStudentRequestAndResponse.ViewStudentRequest{}
	var response = &ViewStudentRequestAndResponse.ViewStudentResponse{}
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
	response.Data.Users, response.Code = MentalListService.MentalListService{}.ViewStudent(cookie)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
}
