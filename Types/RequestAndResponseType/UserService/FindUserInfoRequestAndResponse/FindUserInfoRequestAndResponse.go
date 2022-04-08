package FindUserInfoRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/User"
)

type FindUserInfoRequest struct {
}

type FindUserInfoResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		User     User.User
		UserInfo string `json:"userInfo"`
		Message  string `json:"message"`
	}
}
