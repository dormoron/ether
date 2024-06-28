package users

import (
	"context"
	"errors"
	"ether/internal/domain/model"
	"ether/internal/domain/repository/users"
	"ether/pkg/casbin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
)

var ErrUserDuplicateEmail = users.ErrUserDuplicate
var ErrInvalidUserOrPassword = errors.New("账号/邮箱或密码不对")

type UserService interface {
	Login(ctx context.Context, username, password string) (model.Auth, error)
	SignUp(ctx context.Context, a model.Auth) error
	Profile(ctx context.Context, id uint) (model.User, error)
}

type UserServiceStruct struct {
	authRepo users.AuthRepository
	userRepo users.UserRepository
	roleRepo users.RoleRepository
	casbin.Access
	tx *gorm.DB
}

func NewUserService(authRepo users.AuthRepository, userRepo users.UserRepository,
	roleRepo users.RoleRepository,
	tx *gorm.DB, access casbin.Access) UserService {
	return &UserServiceStruct{
		authRepo: authRepo,
		userRepo: userRepo,
		roleRepo: roleRepo,
		tx:       tx,
		Access:   access,
	}
}

func (svc *UserServiceStruct) Login(ctx context.Context, username, password string) (model.Auth, error) {
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

func (svc *UserServiceStruct) SignUp(ctx context.Context, a model.Auth) error {
	tx := svc.tx.Begin()
	// 查询默认角色
	role := model.Role{
		Name:        "游客",
		Description: "",
		IsDefault:   true,
	}
	rId, err := svc.roleRepo.FirstByDefaultInTxn(ctx, tx, role)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 创建用户
	uId, err := svc.userRepo.CreateInTxn(ctx, tx, a)
	if err != nil {
		tx.Rollback()
		return err
	}
	ok, err := svc.AddRoleForUser(strconv.Itoa(int(uId)), strconv.Itoa(int(rId)))
	if err != nil && !ok {
		tx.Rollback()
		return err
	}
	// 加密
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return err
	}
	a.UserId = uId
	a.Password = string(hash)
	err = svc.authRepo.CreateInTxn(ctx, tx, a)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (svc *UserServiceStruct) Profile(ctx context.Context, id uint) (model.User, error) {
	u, err := svc.userRepo.FindById(ctx, id)
	return u, err
}
