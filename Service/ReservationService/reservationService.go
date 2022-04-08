package ReservationService

import (
	"TheLabSystem/Dao/DeviceDao"
	"TheLabSystem/Dao/ReservationDao"
	"TheLabSystem/Dao/ReservationInfoDao"
	"TheLabSystem/Dao/ReservationRecordDao"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/SubmitReservationRequestAndResponse"
	"TheLabSystem/Types/ServiceType/Device"
	"TheLabSystem/Types/ServiceType/Reservation"
	"TheLabSystem/Types/ServiceType/ReservationInfo"
	"TheLabSystem/Types/ServiceType/ReservationRecord"
	"time"
)

type ReservationService struct {
}

func (service ReservationService) SubmitReservation(username string, request *SubmitReservationRequestAndResponse.SubmitReservationRequest) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	} else if user.Username == "" {
		return ErrNo.LoginRequired
	}
	var reservationDay time.Time
	reservationDay, err = time.Parse("2006-01-02", request.Day)
	if err != nil {
		return ErrNo.ParamInvalid
	}
	if reservationDay.Before(time.Now()) {
		return ErrNo.ParamInvalid
	}
	if request.Time >= 12 || request.Time < 0 {
		return ErrNo.ParamInvalid
	}
	if request.Num <= 0 {
		return ErrNo.ParamInvalid
	}
	// check if there are still enough device
	var devices []Device.Device
	devices, err = DeviceDao.FindDeviceByType(request.DeviceType)
	if len(devices) <= request.Num {
		return ErrNo.ParamInvalid
	}
	reservation := Reservation.Reservation{
		ApplicantID: user.UserID,
	}
	if user.UserType == 1 {
		reservation.Status = 21234
	} else if user.UserType == 2 {
		reservation.Status = 112
	} else if user.UserType == 3 {
		reservation.Status = 32
	}
	reservation, err = ReservationDao.InsertReservation(reservation)
	if err != nil {
		return ErrNo.UnknownError
	}
	var record = ReservationRecord.ReservationRecord{
		ReservationID: reservation.ReservationID,
		OperatorID:    user.UserID,
		OperationType: 1,
		OperatingDay:  reservation.OperatingDay,
	}
	if ReservationRecordDao.InsertReservationRecord(record) != nil {
		return ErrNo.UnknownError
	}
	for i := 0; i < request.Num; i++ {
		info := ReservationInfo.ReservationInfo{
			ReservationID:   reservation.ReservationID,
			DeviceID:        devices[i].DeviceID,
			ReservationDay:  request.Day,
			ReservationTime: request.Time,
		}
		if ReservationInfoDao.InsertReservationInfo(info) != nil {
			return ErrNo.UnknownError
		}
	}
	return ErrNo.OK
}
