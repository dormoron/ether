package users

import (
	"ether/internal/domain/repository/mapper"
	"time"
)

type UserInfo struct {
	mapper.BaseModel
	Email     string `gorm:"not null;comment:'email'"`
	WebSite   string `gorm:"comment:'个人网站'"`
	Nickname  string `gorm:"comment:'用户昵称'"`
	Avatar    string `gorm:"comment:'用户头像'"`
	Intro     string `gorm:"comment:'用户简介'"`
	IsDisable bool   `gorm:"comment:'是否禁用'"`
}

type UserAuth struct {
	mapper.BaseModel
	UserId        string    `gorm:"not null;comment:'用户id'"`
	Username      string    `gorm:"comment:'用户名'"`
	Password      string    `gorm:"comment:'密码'"`
	LoginType     string    `gorm:"comment:'登录类型'"`
	IpAddress     string    `gorm:"comment:'用户登录ip'"`
	IpSource      string    `gorm:"comment:'ip来源'"`
	LastLoginTime time.Time `gorm:"comment:'上次登录时间'"`
}

type Role struct {
	mapper.BaseModel
	Name        string `gorm:"comment:'角色名'"`
	Description string `gorm:"comment:'角色描述'"`
	IsDisable   bool   `gorm:"comment:'是否禁用'"`
}

type UserRole struct {
	mapper.BaseModel
	UserId string `gorm:"not null;comment:'用户id'"`
	RoleId string `gorm:"not null;comment:'角色id'"`
}
