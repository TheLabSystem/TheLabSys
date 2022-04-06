package RegisterUserRequestAndResponse

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}
type RegisterUserResponse struct {
	Code error
}
