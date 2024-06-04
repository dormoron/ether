package users

type SignUpReq struct {
	Username        string `json:"username"`
	ConfirmPassword string `json:"confirmPassword"`
	Password        string `json:"password"`
}
