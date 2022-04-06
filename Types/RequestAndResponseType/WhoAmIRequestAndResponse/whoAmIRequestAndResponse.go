package WhoAmIRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/User"
)

type WhoAmIRequest struct {
}

type WhoAmIResponse struct {
	Code ErrNo.ErrNo
	Data struct {
		User User.User
	}
}
