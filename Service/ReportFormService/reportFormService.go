package ReportFormService

import (
	"TheLabSystem/Config/UserPermissionDecide"
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/ReportForm"
	"fmt"
	"time"
)

type ReportFormService struct {
}

func (service ReportFormService) GetReportForm(form *ReportForm.ReportForm, username string) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username == "" {
		return ErrNo.LoginRequired
	}
	if UserPermissionDecide.GetReportForm(user.UserType) {
		var tm1 time.Time
		var tm2 time.Time
		var err error
		if tm1, err = time.Parse("2006-01-12", form.StartDay); err != nil {
			return ErrNo.ParamInvalid
		}
		if tm2, err = time.Parse("2006-01-12", form.EndDay); err != nil {
			return ErrNo.ParamInvalid
		}

		fmt.Println(tm1, tm2)
		return ErrNo.OK
	} else {
		return ErrNo.PermDenied
	}
}
