package users

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Username        string `json:"username"`
	ConfirmPassword string `json:"confirmPassword"`
	Password        string `json:"password"`
}

type RoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type IdRequest struct {
	Id uint `json:"id"`
}

type AccessRequest struct {
	Subject string `json:"subject"`
	Object  string `json:"object"`
	Action  string `json:"action"`
}
