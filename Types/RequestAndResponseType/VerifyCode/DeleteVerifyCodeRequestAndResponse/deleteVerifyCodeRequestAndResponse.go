package DeleteVerifyCodeRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type DeleteVerifyCodeRequest struct {
	VerifyCode int `json:"deleteVerifyCode"`
}

type DeleteVerifyCodeResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
