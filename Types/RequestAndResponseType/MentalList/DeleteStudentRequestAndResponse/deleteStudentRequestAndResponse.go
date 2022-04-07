package DeleteStudentRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type DeleteStudentRequest struct {
	StudentID uint `json:"student_id"`
}

type DeleteStudentResponse struct {
	Code ErrNo.ErrNo `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
