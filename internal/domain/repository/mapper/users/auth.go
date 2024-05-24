package users

import "gorm.io/gorm"

type UserAuthMapper struct {
	db *gorm.DB
}

func NewUserAuthMapper(db *gorm.DB) *UserAuthMapper {
	return &UserAuthMapper{
		db: db,
	}
}
