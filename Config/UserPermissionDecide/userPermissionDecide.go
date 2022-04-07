package UserPermissionDecide

func AddVerifyCode(UserType int) bool {
	return UserType >= 4
}

func ViewVerifyCode(UserType int) bool {
	return UserType >= 4
}

func DeleteVerifyCode(UserType int) bool {
	return UserType >= 4
}
