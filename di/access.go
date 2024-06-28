package di

import (
	"ether/pkg/casbin"
	"gorm.io/gorm"
)

func InitAccess(db *gorm.DB) casbin.Access {
	accessCasbin, err := casbin.NewAccessCasbin(db, "config/model.conf")
	if err != nil {
		panic(err)
	}
	return accessCasbin
}
