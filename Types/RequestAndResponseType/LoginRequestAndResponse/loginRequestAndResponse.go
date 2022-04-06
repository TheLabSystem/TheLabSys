package LoginRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/User"
)

type LoginRequest struct {
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
type LoginResponse struct {
	Code ErrNo.ErrNo `json:"ErrCode" binding:"required"`
	Data struct {
		Message  string `json:"message" binding:"required"`
		User     User.User
		Userinfo string `json:"userinfo" `
	}
}
