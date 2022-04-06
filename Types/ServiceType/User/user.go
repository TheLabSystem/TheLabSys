package User

type User struct {
	UserID      uint    `json:"user_id"`
	Username    string  `json:"user_name"`
	UserType    int     `json:"user-type"`
	DisplayName string  `json:"display_name"`
	Password    string  `json:"password"`
	Money       float64 `json:"money"`
}
