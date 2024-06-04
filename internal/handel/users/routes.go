package users

import "github.com/dormoron/mist"

func (u *AuthHandler) RegisterRoutes(server mist.HTTPServer) {
	userGroup := server.Group("users")
	userGroup.POST("/signup", u.SignUp)
}
