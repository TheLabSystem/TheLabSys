package ReservationController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Service/ReservationService"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/RevertReservationRequestAndRespoonse"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/SubmitReservationRequestAndResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReservationController struct {
}

func (controller ReservationController) SubmitReservation(c *gin.Context) {
	var request = &SubmitReservationRequestAndResponse.SubmitReservationRequest{}
	var response = &SubmitReservationRequestAndResponse.SubmitReservationResponse{}
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
	response.Code = ReservationService.ReservationService{}.SubmitReservation(cookie, request)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
	return
}

func (controller ReservationController) RevertReservation(c *gin.Context) {
	var request = &RevertReservationRequestAndRespoonse.RevertReservationRequest{}
	var response = &RevertReservationRequestAndRespoonse.RevertReservationResponse{}
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
	response.Code = ReservationService.ReservationService{}.RevertReservation(cookie, request.ReservationID)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
	return
}
