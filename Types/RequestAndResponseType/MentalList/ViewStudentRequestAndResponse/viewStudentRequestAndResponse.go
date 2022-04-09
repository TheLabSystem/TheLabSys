package ViewStudentRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/User"
)

type ViewStudentRequest struct {
}
type ViewStudentResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Users   []User.User `json:"users"`
		Message string      `json:"message"`
	}
}
