package users

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UserAuthMapper struct {
	db *gorm.DB
}

func NewUserAuthMapper(db *gorm.DB) *UserAuthMapper {
	return &UserAuthMapper{
		db: db,
	}
}

func (mapper *UserAuthMapper) Insert(ctx context.Context, a UserAuth) error {
	err := mapper.db.WithContext(ctx).Create(&a).Error
	var me *mysql.MySQLError
	if errors.As(err, &me) {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			return ErrDuplicateEmail
		}
	}
	return err
}

func (mapper *UserAuthMapper) Delete(ctx context.Context, id string) error {
	return mapper.db.WithContext(ctx).Where("id = ?", id).Delete(&UserAuth{}).Error
}
