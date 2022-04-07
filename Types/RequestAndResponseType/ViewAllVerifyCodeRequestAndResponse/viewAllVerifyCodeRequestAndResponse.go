package ViewAllVerifyCodeRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/VerifyCode"
)

type ViewAllVerifyCodeRequest struct {
}

type ViewAllVerifyCodeResponse struct {
	Code ErrNo.ErrNo
	Data struct {
		Message     string                  `json:"message"`
		VerifyCodes []VerifyCode.VerifyCode `json:"verify_codes"`
	}
}
