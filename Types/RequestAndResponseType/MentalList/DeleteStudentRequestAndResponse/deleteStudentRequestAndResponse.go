package DeleteStudentRequestAndResponse

type DeleteStudentRequest struct {
	StudentID uint
}

type DeleteStudentResponse struct {
	Code int `json:"Code"`
	Data struct {
		Message string `json:"message"`
	}
}
