package BillService

import (
	"TheLabSystem/Dao/BillDao"
	"TheLabSystem/Dao/ReservationDao"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/Bill"
)

type BillService struct {
}

func (billService BillService) GetBill(username string) ([]Bill.Bill, ErrNo.ErrNo) {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return nil, ErrNo.UnknownError
	}
	if user.Username == "" {
		return nil, ErrNo.LoginRequired
	}
	bill, errB := BillDao.FindBillsByPayerID(user.UserID)
	if errB != nil {
		return nil, ErrNo.UnknownError
	}
	return bill, ErrNo.OK
}

func (billService BillService) PayBill(billID uint, username string) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username == "" {
		return ErrNo.LoginRequired
	}
	bill, err := BillDao.FindBillByBillID(billID)
	if err != nil {
		return ErrNo.UnknownError
	}
	if bill.Money > user.Money {
		return ErrNo.MoneyNotEnough
	} else {
		if err := UserDao.UpdateMoney(user, user.Money-bill.Money); err != nil {
			return ErrNo.UnknownError
		}
		if err := BillDao.UpdateBillStatus(billID, 1); err != nil {
			return ErrNo.UnknownError
		}
		if reservation, err := ReservationDao.FindReservationByID(bill.ReservationID); err != nil {
			return ErrNo.UnknownError
		} else {
			if reservation.Status == 2234 {
				if err := ReservationDao.UpdateReservation(reservation.ReservationID, 234); err != nil {
					return ErrNo.UnknownError
				}
			}
		}
	}
	return ErrNo.OK
}
