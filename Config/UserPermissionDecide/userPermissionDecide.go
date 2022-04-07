package UserPermissionDecide

func AddVerifyCode(UserType int) bool {
	return UserType >= 4
}
