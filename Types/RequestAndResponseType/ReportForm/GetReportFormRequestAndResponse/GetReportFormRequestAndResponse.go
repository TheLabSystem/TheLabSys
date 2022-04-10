package GetReportFormRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/ReportForm"
)

type GetReportFormRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
type GetReportFormResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		ReportForm ReportForm.ReportForm `json:"reportForm"`
		Message    string                `json:"message"`
	}
}
