package ReservationService

import (
	"TheLabSystem/Config/UserPermissionDecide"
	"TheLabSystem/Dao/BillDao"
	"TheLabSystem/Dao/DeviceDao"
	"TheLabSystem/Dao/DeviceTypeInfoDao"
	"TheLabSystem/Dao/MentorRecordDao"
	"TheLabSystem/Dao/ReservationDao"
	"TheLabSystem/Dao/ReservationInfoDao"
	"TheLabSystem/Dao/ReservationRecordDao"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/GetApprovalRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/SetApprovalRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/SubmitReservationRequestAndResponse"
	"TheLabSystem/Types/ServiceType/Bill"
	"TheLabSystem/Types/ServiceType/Device"
	"TheLabSystem/Types/ServiceType/DeviceTypeInfo"
	"TheLabSystem/Types/ServiceType/Reservation"
	"TheLabSystem/Types/ServiceType/ReservationInfo"
	"TheLabSystem/Types/ServiceType/ReservationRecord"
	"TheLabSystem/Types/ServiceType/User"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"time"
)

type Approval struct {
	reservation Reservation.Reservation
	user        User.User
}

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
	var deviceTypeInfo DeviceTypeInfo.DeviceTypeInfo
	deviceTypeInfo, err = DeviceTypeInfoDao.FindDeviceTypeInfoByDeviceTypeID(request.DeviceType)
	if err != nil {
		return ErrNo.UnknownError
	}
	if len(devices) < request.Num {
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
			DeviceTypeInfo:  deviceTypeInfo.DeviceInfo,
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
	if user.UserType == 1 {
		var bill Bill.Bill
		bill, err = BillDao.FindBillByReservationID(reservationID)
		if err != nil {
			return ErrNo.UnknownError
		}
		if bill.BillStatus == 1 {
			user.Money += 0.95 * bill.Money
			err = BillDao.UpdateBillStatus(bill.BillID, -1)
			if err != nil {
				return ErrNo.UnknownError
			}
			err = UserDao.UpdateUser(user)
			if err != nil {
				return ErrNo.UnknownError
			}
		}
	}

	return ErrNo.OK
}
func (service ReservationService) GetApproval(username string) ([]GetApprovalRequestAndResponse.Approval, ErrNo.ErrNo) {
	var approval []GetApprovalRequestAndResponse.Approval
	var reservation []Reservation.Reservation
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return approval, ErrNo.UnknownError
	} else if user.Username == "" {
		return approval, ErrNo.LoginRequired
	}
	mentorRecord, err := MentorRecordDao.FindMentorRecordByTeacherID(user.UserID)
	if err != nil {
		return approval, ErrNo.UnknownError
	}
	reservation, err = ReservationDao.FindApprovalReservation(user.UserType, mentorRecord)
	if err != nil {
		return approval, ErrNo.UnknownError
	}
	approval = make([]GetApprovalRequestAndResponse.Approval, len(reservation), len(reservation))
	for key := range reservation {
		approval[key].ReservationRes = reservation[key]
		approval[key].UserRes, err = UserDao.FindUserByID(reservation[key].ApplicantID)
		if err != nil {
			return nil, ErrNo.UnknownError
		}
	}
	return approval, ErrNo.OK
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
				if err := ReservationDao.UpdateReservation(request.ReservationID, -2); err != nil {
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
				if err := ReservationDao.UpdateReservation(request.ReservationID, -2); err != nil {
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
				reservationInfo, err := ReservationInfoDao.FindInfoByReservationID(request.ReservationID)
				if err != nil {
					return ErrNo.UnknownError
				}
				device, err := DeviceDao.FindDeviceByDeviceID(reservationInfo[0].DeviceID)
				if err != nil {
					return ErrNo.UnknownError
				}
				deviceInfo, err := DeviceTypeInfoDao.FindDeviceTypeInfoByDeviceTypeID(device.DeviceTypeID)
				if err != nil {
					return ErrNo.UnknownError
				}
				if err := BillDao.InsertBill(Bill.Bill{
					ReservationID: request.ReservationID,
					PayerID:       reservation.ApplicantID,
					Money:         deviceInfo.Money * float64(len(reservationInfo)),
					BillStatus:    2,
				}); err != nil {
					return ErrNo.UnknownError
				}
			} else if request.Approval == 2 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, -2); err != nil {
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
				if err := ReservationDao.UpdateReservation(request.ReservationID, -2); err != nil {
					return ErrNo.UnknownError
				}
			}
		} else {
			return ErrNo.PermDenied
		}
	}
	if reservation.Status == 24 {
		if user.UserType == 4 {
			if request.Approval == 1 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, 2); err != nil {
					return ErrNo.UnknownError
				}
			} else if request.Approval == 2 {
				if err := ReservationDao.UpdateReservation(request.ReservationID, -2); err != nil {
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
				if err := ReservationDao.UpdateReservation(request.ReservationID, -2); err != nil {
					return ErrNo.UnknownError
				}
			}
		} else {
			return ErrNo.PermDenied
		}
	}
	if request.Approval == 1 {
		if err := ReservationRecordDao.InsertReservationRecord(ReservationRecord.ReservationRecord{
			ReservationID: request.ReservationID,
			OperatorID:    user.UserID,
			OperationType: 2,
		}); err != nil {
			return ErrNo.UnknownError
		}
	}
	if request.Approval == 2 {
		if err := ReservationRecordDao.InsertReservationRecord(ReservationRecord.ReservationRecord{
			ReservationID: request.ReservationID,
			OperatorID:    user.UserID,
			OperationType: 3,
		}); err != nil {
			return ErrNo.UnknownError
		}
	}
	return ErrNo.OK
}
func (service ReservationService) GetPersonalReservations(username string) ([]Reservation.Reservation, ErrNo.ErrNo) {
	user, err := UserDao.FindUserByUsername(username)
	var res []Reservation.Reservation
	if err != nil {
		return res, ErrNo.UnknownError
	} else if user.Username == "" {
		return res, ErrNo.LoginRequired
	}
	res, err = ReservationDao.FindReservationByApplicantID(user.UserID)
	if err != nil {
		return res, ErrNo.UnknownError
	} else {
		return res, ErrNo.OK
	}
}
func (service ReservationService) GetReservationByReservationID(username string, reservationID uint) (Reservation.Reservation, []ReservationInfo.ReservationInfo, ErrNo.ErrNo) {
	user, err := UserDao.FindUserByUsername(username)
	var res []ReservationInfo.ReservationInfo
	var res1 Reservation.Reservation
	if err != nil {
		return res1, nil, ErrNo.UnknownError
	} else if user.Username == "" {
		return res1, nil, ErrNo.LoginRequired
	}
	res, err = ReservationInfoDao.FindInfoByReservationID(reservationID)
	if err != nil {
		return res1, nil, ErrNo.UnknownError
	}
	res1, err = ReservationDao.FindReservationByID(reservationID)
	if err != nil {
		return res1, nil, ErrNo.UnknownError
	}

	return res1, res, ErrNo.OK
}
func (service ReservationService) GetReservationDetails(username string, day string, deviceTypeID uint) ([]int, ErrNo.ErrNo) {
	user, err := UserDao.FindUserByUsername(username)
	var res []int
	if err != nil {
		return res, ErrNo.UnknownError
	} else if user.Username == "" {
		return res, ErrNo.LoginRequired
	}
	var devices []Device.Device
	devices, err = DeviceDao.FindDeviceByTypeAllowRecordNotFound(deviceTypeID)
	var deviceIdSet = mapset.NewSet()
	for key := range devices {
		if devices[key].DeviceStatus == 2 {
			deviceIdSet.Add(devices[key].DeviceID)
		}
	}
	fmt.Println("getting devices", devices)
	if err != nil {
		return res, ErrNo.UnknownError
	}
	var info []ReservationInfo.ReservationInfo
	info, err = ReservationInfoDao.FindAllReservationInfo()
	if err != nil {
		return res, ErrNo.UnknownError
	}
	res = make([]int, 12, 12)
	for key := range res {
		res[key] = len(deviceIdSet.ToSlice())
	}
	for key := range info {
		if deviceIdSet.Contains(info[key].DeviceID) && info[key].ReservationDay == day {
			var reservation Reservation.Reservation
			reservation, err = ReservationDao.FindReservationByID(info[key].ReservationID)
			if reservation.Status >= 10 || reservation.Status < 0 {
				continue
			}
			if err != nil {
				return nil, ErrNo.UnknownError
			} else {
				if reservation.Status < 10 {
					res[info[key].ReservationTime]--
				}
			}
		}
	}
	return res, ErrNo.OK
}
