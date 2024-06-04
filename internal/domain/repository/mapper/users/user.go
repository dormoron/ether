package users

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrUpdateData     = errors.New("更新数据失败")
	ErrUserNotFound   = gorm.ErrRecordNotFound
)

type UserMapper struct {
	db *gorm.DB
}

func NewUserMapper(db *gorm.DB) *UserMapper {
	return &UserMapper{
		db: db,
	}
}

func (mapper *UserMapper) Insert(ctx context.Context, u User) (int64, error) {
	err := mapper.db.WithContext(ctx).Create(&u).Error
	var me *mysql.MySQLError
	if errors.As(err, &me) {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			return 0, ErrDuplicateEmail
		}
	}
	return u.Id, err
}

func (mapper *UserMapper) UpdateById(ctx context.Context, u User) error {
	res := mapper.db.Model(&u).WithContext(ctx).
		Where("id =?", u.Id).
		Updates(map[string]any{
			"email":    u.Email,
			"webSite":  u.WebSite,
			"nickname": u.Nickname,
			"avatar":   u.Avatar,
			"intro":    u.Intro,
		})
	err := res.Error
	if err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		return ErrUpdateData
	}
	return nil
}

func (mapper *UserMapper) Disable(ctx context.Context, u User) error {
	res := mapper.db.Model(&u).WithContext(ctx).
		Where("id =?", u.Id).
		Updates(map[string]any{
			"isDisable": u.IsDisable,
		})
	err := res.Error
	if err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		return ErrUpdateData
	}
	return nil
}

func (mapper *UserMapper) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := mapper.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

func (mapper *UserMapper) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := mapper.db.WithContext(ctx).Where("id=?", id).First(&u).Error
	return u, err
}

func (mapper *UserMapper) FindByDisable(ctx context.Context, disable bool) ([]User, error) {
	var u []User
	err := mapper.db.WithContext(ctx).Where("is_disable=?", disable).Find(&u).Error
	return u, err
}
