package GetReportFormRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
)

type GetReportFormRequest struct {
}
type GetReportFormResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
