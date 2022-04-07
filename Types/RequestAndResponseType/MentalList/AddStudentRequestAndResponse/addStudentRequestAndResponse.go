package AddStudentRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type AddStudentRequest struct {
	StudentID uint `json:"student_id"`
}
type AddStudentResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
