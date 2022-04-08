package ViewStudentRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/MentorRecord"
)

type ViewStudentRequest struct {
}
type ViewStudentResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		MentorRecords []MentorRecord.MentorRecord `json:"mentorRecords"`
		Message       string                      `json:"message"`
	}
}
