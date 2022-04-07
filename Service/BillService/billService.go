package BillService

import (
	"TheLabSystem/Dao/BillDao"
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
