package RegisterUserRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type RegisterUserRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	UserType   int    `json:"user_type"`
	VerifyCode int    `json:"verify_code"`
}
type RegisterUserResponse struct {
	Code ErrNo.ErrNo
	Data struct {
		Message string `json:"message"`
	}
}
