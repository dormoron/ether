package users

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Auth struct {
	gorm.Model
	UserId        uint           `gorm:"unique;comment:'用户id'"`
	Username      sql.NullString `gorm:"unique;comment:'用户名'"`
	Password      string         `gorm:"comment:'密码'"`
	LoginType     string         `gorm:"comment:'登录类型'"`
	IpAddress     string         `gorm:"comment:'用户登录ip'"`
	IpSource      string         `gorm:"comment:'ip来源'"`
	LastLoginTime time.Time      `gorm:"comment:'上次登录时间'"`
}

type User struct {
	gorm.Model
	Email     sql.NullString `gorm:"unique;comment:'email'"`
	WebSite   string         `gorm:"comment:'个人网站'"`
	Nickname  string         `gorm:"comment:'用户昵称'"`
	Avatar    string         `gorm:"comment:'用户头像'"`
	Intro     string         `gorm:"comment:'用户简介'"`
	IsDisable bool           `gorm:"comment:'是否禁用'"`
}

type Role struct {
	gorm.Model
	Name        string `gorm:"unique;comment:'角色名'"`
	Description string `gorm:"comment:'角色描述'"`
	IsDefault   bool   `gorm:"comment:'是否默认'"`
	IsDisable   bool   `gorm:"comment:'是否禁用'"`
}
