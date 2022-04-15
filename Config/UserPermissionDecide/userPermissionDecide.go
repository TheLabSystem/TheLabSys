package UserPermissionDecide

func AddVerifyCode(UserType int) bool {
	return UserType == 4 || UserType == 255
}

func ViewVerifyCode(UserType int) bool {
	return UserType >= 4
}

func DeleteVerifyCode(UserType int) bool {
	return UserType == 4 || UserType == 255
}

func SetApproval(UserType int) bool {
	return UserType >= 3
}

func GetReportForm(UserType int) bool {
	return UserType >= 5
}

func AddDevice(UserType int) bool {
	return UserType == 255
}
func UpdateDevice(UserType int) bool {
	return UserType == 4 || UserType == 255
}
