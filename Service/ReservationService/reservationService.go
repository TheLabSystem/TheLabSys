package ReservationService

import (
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	SubmitReservationAndResponse "TheLabSystem/Types/RequestAndResponseType/Reservation"
	"time"
)

type ReservationService struct {
}

func (service ReservationService) SubmitReservation(username string, request SubmitReservationAndResponse.SubmitReservationRequest) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	} else if user.Username == "" {
		return ErrNo.LoginRequired
	}
	var reservation_day time.Time
	reservation_day, err = time.Parse("2006-01-02", request.Day)
	if err != nil {
		return ErrNo.ParamInvalid
	}
	if reservation_day.Before(time.Now()) {
		return ErrNo.ParamInvalid
	}
	if request.Time >= 12 || request.Time < 0 {
		return ErrNo.ParamInvalid
	}
	return ErrNo.OK
}
