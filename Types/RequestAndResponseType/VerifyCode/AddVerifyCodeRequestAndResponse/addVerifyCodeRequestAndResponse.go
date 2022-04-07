package AddVerifyCodeRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type AddVerifyCodeRequest struct {
	VerifyCode int `json:"verify_code"`
	UserType   int `json:"user_type"`
}
type AddVerifyCodeResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
