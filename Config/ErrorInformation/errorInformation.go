package ErrorInformation

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

func GenerateErrorInformation(code ErrNo.ErrNo) string {
	if code == ErrNo.UnknownError {
		return "An unexpected error has happened.Please try again."
	}
	if code == ErrNo.WrongPassword {
		return "Wrong Password!"
	}
	if code == ErrNo.LoginRequired {
		return "Your information is out of date.Please log in first."
	}
	if code == ErrNo.UserNotExisted {
		return "There is no such user.Maybe your account has been deleted,so try connecting the administrator."
	}
	if code == ErrNo.UserHasExisted {
		return "This username has been registered before."
	}
	if code == ErrNo.PermDenied {
		return "Sorry.You don't have permission to do so."
	}
	if code == ErrNo.VerifyCodeNotValid {
		return "It seems that the verify code is not valid now.Try connecting the administrator."
	}
	if code == ErrNo.StudentNotExist {
		return "The student doesn't exist.Please check and try again."
	}
	return "Success!"
}
