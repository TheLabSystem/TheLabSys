package ReservationService

import (
	"TheLabSystem/Config/UserPermissionDecide"
	"TheLabSystem/Dao/DeviceDao"
	"TheLabSystem/Dao/ReservationDao"
	"TheLabSystem/Dao/ReservationInfoDao"
	"TheLabSystem/Dao/ReservationRecordDao"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/GetApprovalRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/SetApprovalRequestAndResponse"
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

func (service ReservationService) RevertReservation(username string, reservationID uint) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	} else if user.Username == "" {
		return ErrNo.LoginRequired
	}
	if ReservationDao.UpdateReservation(reservationID, -1) != nil {
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
func (service ReservationService) GetApproval(username string, request *GetApprovalRequestAndResponse.GetApprovalRequest) ([]Reservation.Reservation, ErrNo.ErrNo) {
	var reservation []Reservation.Reservation
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return reservation, ErrNo.UnknownError
	} else if user.Username == "" {
		return reservation, ErrNo.LoginRequired
	}
	if request.Status == 1 {
		reservation, err = ReservationDao.FindAllReservation()
		if err != nil {
			return reservation, ErrNo.UnknownError
		}
	} else if request.Status == 2 {
		reservation, err = ReservationDao.FindApprovalReservation()
		if err != nil {
			return reservation, ErrNo.UnknownError
		}
	} else if request.Status == 3 {
		reservation, err = ReservationDao.FindDisapprovalReservation()
		if err != nil {
			return reservation, ErrNo.UnknownError
		}
	}
	return reservation, ErrNo.OK
}
func (service ReservationService) SetApproval(username string, request *SetApprovalRequestAndResponse.SetApprovalRequest) ErrNo.ErrNo {
	reservation, errB := ReservationDao.FindReservationByID(request.ReservationID)
	if errB != nil {
		return ErrNo.UnknownError
	}
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	} else if user.Username == "" {
		return ErrNo.LoginRequired
	} else if !UserPermissionDecide.SetApproval(user.UserType) {
		return ErrNo.PermDenied
	}
	if reservation.Status == 112 {
		if user.UserType == 3 {
			if request.Approval == 1 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, 12); err != nil {
					return ErrNo.UnknownError
				}
			} else if request.Approval == 2 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, -1); err != nil {
					return ErrNo.UnknownError
				}
			}
		} else {
			return ErrNo.PermDenied
		}
	}
	if reservation.Status == 12 {
		if user.UserType == 4 {
			if request.Approval == 1 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, 1); err != nil {
					return ErrNo.UnknownError
				}
			} else if request.Approval == 2 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, -1); err != nil {
					return ErrNo.UnknownError
				}
			}
		} else {
			return ErrNo.PermDenied
		}
	}
	if reservation.Status == 21234 {
		if user.UserType == 4 {
			if request.Approval == 1 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, 2234); err != nil {
					return ErrNo.UnknownError
				}
			} else if request.Approval == 2 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, -1); err != nil {
					return ErrNo.UnknownError
				}
			}
		} else {
			return ErrNo.PermDenied
		}
	}
	if reservation.Status == 234 {
		if user.UserType == 255 {
			if request.Approval == 1 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, 24); err != nil {
					return ErrNo.UnknownError
				}
			} else if request.Approval == 2 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, -1); err != nil {
					return ErrNo.UnknownError
				}
			}
		} else {
			return ErrNo.PermDenied
		}
	}
	if reservation.Status == 24 {
		if user.UserType == 255 {
			if request.Approval == 1 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, 2); err != nil {
					return ErrNo.UnknownError
				}
			} else if request.Approval == 2 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, -1); err != nil {
					return ErrNo.UnknownError
				}
			}
		} else {
			return ErrNo.PermDenied
		}
	}
	if reservation.Status == 32 {
		if user.UserType == 4 {
			if request.Approval == 1 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, 3); err != nil {
					return ErrNo.UnknownError
				}
			} else if request.Approval == 2 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, -1); err != nil {
					return ErrNo.UnknownError
				}
			}
		} else {
			return ErrNo.PermDenied
		}
	}
	return ErrNo.OK
}
