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

type UserInfoMapper struct {
	db *gorm.DB
}

func NewUserInfoMapper(db *gorm.DB) *UserInfoMapper {
	return &UserInfoMapper{
		db: db,
	}
}

func (mapper *UserInfoMapper) Insert(ctx context.Context, u Info) error {
	err := mapper.db.WithContext(ctx).Create(&u).Error
	var me *mysql.MySQLError
	if errors.As(err, &me) {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			return ErrDuplicateEmail
		}
	}
	return err
}

func (mapper *UserInfoMapper) UpdateById(ctx context.Context, u Info) error {
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

func (mapper *UserInfoMapper) Disable(ctx context.Context, u Info) error {
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

func (mapper *UserInfoMapper) FindByEmail(ctx context.Context, email string) (Info, error) {
	var u Info
	err := mapper.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

func (mapper *UserInfoMapper) FindByDisable(ctx context.Context, disable bool) ([]Info, error) {
	var u []Info
	err := mapper.db.WithContext(ctx).Where("is_disable=?", disable).Find(&u).Error
	return u, err
}
