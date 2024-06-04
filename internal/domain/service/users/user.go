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

type UserService struct {
	authRepo users.AuthRepository
	userRepo users.UserRepository
}

func NewUserService(authRepo users.AuthRepository, userRepo users.UserRepository) *UserService {
	return &UserService{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (svc *UserService) Login(ctx context.Context, username, password string) (model.Auth, error) {
	// 判断账户有没有禁用
	u, err := svc.userRepo.FindByEmail(ctx, username)
	if err != nil {
		return model.Auth{}, err
	}
	if u.IsDisable {
		return model.Auth{}, errors.New("用户已被禁用")
	}
	// 登录
	a, err := svc.authRepo.FindByUsername(ctx, username)
	if errors.Is(err, users.ErrUserNotFound) {
		return model.Auth{}, ErrInvalidUserOrPassword
	}
	err = bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	if err != nil {
		return model.Auth{}, ErrInvalidUserOrPassword
	}
	return a, nil
}

func (svc *UserService) SignUp(ctx context.Context, a model.Auth) error {
	// 创建用户
	uId, err := svc.userRepo.Create(ctx, a)
	if err != nil {
		return err
	}
	// 加密
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.UserId = uId
	a.Password = string(hash)
	return svc.authRepo.Create(ctx, a)
}
