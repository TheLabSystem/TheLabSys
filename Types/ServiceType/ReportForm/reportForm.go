package ReportForm

type ReportForm struct {
	StartDay                         string  `json:"startDay"` // for example "2020-12-9"
	EndDay                           string  `json:"endDay"`
	ForeignUserReservation           int     `json:"foreignUserReservation"`
	SuccessfulForeignUserReservation int     `json:"successfulForeignUserReservation"`
	StudentReservation               int     `json:"studentReservation"`
	SuccessfulStudentReservation     int     `json:"successfulStudentReservation"`
	TeacherReservation               int     `json:"teacherReservation"`
	SuccessfulTeacherReservation     int     `json:"successfulTeacherReservation"`
	MoneyIn                          float64 `json:"moneyIn"`
	MoneyOut                         float64 `json:"moneyOut"`
}
