package users

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type AuthMapper interface {
	InsertInTxn(ctx context.Context, tx *gorm.DB, a Auth) error
	FindByUsername(ctx context.Context, username string) (Auth, error)
	Delete(ctx context.Context, id uint) error
}

type AuthMapperStruct struct {
	db *gorm.DB
}

func NewUserAuthMapper(db *gorm.DB) AuthMapper {
	return &AuthMapperStruct{
		db: db,
	}
}

func (mapper *AuthMapperStruct) InsertInTxn(ctx context.Context, tx *gorm.DB, a Auth) error {
	err := tx.WithContext(ctx).Create(&a).Error
	var me *mysql.MySQLError
	if errors.As(err, &me) {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			return ErrDuplicateEmail
		}
	}
	return err
}

func (mapper *AuthMapperStruct) FindByUsername(ctx context.Context, username string) (Auth, error) {
	var u Auth
	err := mapper.db.WithContext(ctx).Where("username = ?", username).First(&u).Error
	return u, err
}

func (mapper *AuthMapperStruct) Delete(ctx context.Context, id uint) error {
	return mapper.db.WithContext(ctx).Where("id = ?", id).Delete(&Auth{}).Error
}
