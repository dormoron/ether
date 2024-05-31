package users

import (
	"context"
	"errors"
	"ether/internal/domain/model"
	"ether/internal/domain/repository/users"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicateEmail = users.ErrUserDuplicate
var ErrInvalidUserOrPassword = errors.New("账号/邮箱或密码不对")

type AuthService struct {
	repo users.AuthRepository
}

func NewAuthService(repo users.AuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (svc *AuthService) Login(ctx context.Context, username, password string) (model.Auth, error) {
	u, err := svc.repo.FindByUsername(ctx, username)
	if errors.Is(err, users.ErrUserNotFound) {
		return model.Auth{}, ErrInvalidUserOrPassword
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return model.Auth{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *AuthService) SignUp(ctx context.Context, u model.Auth) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}
