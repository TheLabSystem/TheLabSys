package ReportFormService

import (
	"TheLabSystem/Config/UserPermissionDecide"
	"TheLabSystem/Dao/BillDao"
	"TheLabSystem/Dao/ReservationDao"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/Bill"
	"TheLabSystem/Types/ServiceType/ReportForm"
	"TheLabSystem/Types/ServiceType/Reservation"
	"TheLabSystem/Types/ServiceType/User"
	"fmt"
	"time"
)

type ReportFormService struct {
}

func (service ReportFormService) GetReportForm(startDay string, endDay string, username string) (ReportForm.ReportForm, ErrNo.ErrNo) {
	var res ReportForm.ReportForm
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return res, ErrNo.UnknownError
	}
	if user.Username == "" {
		return res, ErrNo.LoginRequired
	}
	if UserPermissionDecide.GetReportForm(user.UserType) {
		var tm1 time.Time
		var tm2 time.Time
		var err error
		if tm1, err = time.Parse("2006-01-02", startDay); err != nil {
			fmt.Println(err)
			return res, ErrNo.ParamInvalid
		}
		if tm2, err = time.Parse("2006-01-02", endDay); err != nil {
			fmt.Println(err)
			return res, ErrNo.ParamInvalid
		}
		res.StartDay = startDay
		res.EndDay = endDay
		var reservations []Reservation.Reservation
		reservations, err = ReservationDao.FindAllReservation()
		if err != nil {
			fmt.Println(err)
			return res, ErrNo.UnknownError
		}
		for key := range reservations {
			var reservation = reservations[key]
			thisOperationDay, err := time.Parse("2006-01-02", reservation.OperatingDay)
			if err != nil {
				fmt.Println(err)
				return res, ErrNo.UnknownError
			}
			if thisOperationDay.Before(tm1) || tm2.Before(thisOperationDay) {
				continue
			}
			var applicant User.User
			applicant, err = UserDao.FindUserByID(reservation.ApplicantID)
			if err != nil {
				fmt.Println(err)
				return res, ErrNo.UnknownError
			}
			if applicant.UserType == 1 {
				res.ForeignUserReservation++
				if reservation.Status < 10 && reservation.Status > 0 {
					res.SuccessfulForeignUserReservation++
				}
				var bill Bill.Bill
				bill, err = BillDao.FindBillByReservationID(reservation.ReservationID)
				if err != nil {
					fmt.Println(err)
					return res, ErrNo.UnknownError
				}
				if bill.BillStatus == 1 {
					res.MoneyIn += bill.Money
				} else {
					res.MoneyOut -= bill.Money * 0.95
				}
			} else if applicant.UserType == 2 {
				res.StudentReservation++
				if reservation.Status < 10 && reservation.Status > 0 {
					res.SuccessfulStudentReservation++
				}
			} else if applicant.UserType == 3 {
				res.TeacherReservation++
				if reservation.Status < 10 && reservation.Status > 0 {
					res.SuccessfulTeacherReservation++
				}
			}
		}
		return res, ErrNo.OK
	} else {
		return res, ErrNo.PermDenied
	}
}
