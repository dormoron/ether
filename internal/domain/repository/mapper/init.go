package mapper

import (
	"ether/internal/domain/repository/mapper/users"
	"gorm.io/gorm"
)

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(
		&users.UserInfo{},
		&users.UserAuth{},
		&users.Role{},
		&users.UserRole{})
}
