package DeviceController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Service/DeviceService"
	"TheLabSystem/Types/RequestAndResponseType/Device/AddDeviceRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Device/DeleteDeviceRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Device/GetDeviceTypeRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Device/UpdateDeviceRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeviceController struct {
}

func (controller DeviceController) AddDevice(c *gin.Context) {
	var request = &AddDeviceRequestAndResponse.AddDeviceRequest{}
	var response = &AddDeviceRequestAndResponse.AddDeviceResponse{}
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
	response.Code = DeviceService.DeviceService{}.AddDevice(cookie, request.DeviceID, request.DeviceTypeID, request.DeviceInfo)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
	return
}
func (controller DeviceController) GetDeviceType(c *gin.Context) {
	var request = &GetDeviceTypeRequestAndResponse.GetDeviceTypeRequest{}
	var response = &GetDeviceTypeRequestAndResponse.GetDeviceTypeResponse{}
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
	response.Code, response.Data.DeviceTypeInfo = DeviceService.DeviceService{}.GetDeviceType(cookie)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
	return
}
func (controller DeviceController) UpdateDevice(c *gin.Context) {
	var request = &UpdateDeviceRequestAndResponse.UpdateDeviceRequest{}
	var response = &UpdateDeviceRequestAndResponse.UpdateDeviceResponse{}
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
	response.Code = DeviceService.DeviceService{}.UpdateDevice(cookie, request.DeviceID, request.DeviceTypeID, request.DeviceStatus)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
	return
}
func (controller DeviceController) DeleteDevice(c *gin.Context) {
	var request = &DeleteDeviceRequestAndResponse.DeleteDeviceRequest{}
	var response = &DeleteDeviceRequestAndResponse.DeleteDeviceResponse{}
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
	response.Code = DeviceService.DeviceService{}.DeleteDevice(cookie, request.DeviceID)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
	return
}
