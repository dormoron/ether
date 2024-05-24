package users

import (
	"ether/internal/domain/repository/mapper"
	"time"
)

type UserInfo struct {
	mapper.BaseModel
	Email     string `gorm:"unique"`
	WebSite   string `gorm:"comment:'个人网站'"`
	Nickname  string `gorm:"comment:'用户昵称'"`
	Avatar    string `gorm:"comment:'用户头像'"`
	Intro     string `gorm:"comment:'用户简介'"`
	IsDisable bool   `gorm:"comment:'是否禁用'"`
}

type UserAuth struct {
	mapper.BaseModel
	UserId        string    `gorm:"unique;comment:'用户id'"`
	Username      string    `gorm:"comment:'用户名'"`
	Password      string    `gorm:"comment:'密码'"`
	LoginType     string    `gorm:"comment:'登录类型'"`
	IpAddress     string    `gorm:"comment:'用户登录ip'"`
	IpSource      string    `gorm:"comment:'ip来源'"`
	LastLoginTime time.Time `gorm:"comment:'上次登录时间'"`
}

type Permission struct {
	mapper.BaseModel
	Name          string `gorm:"unique;comment:'权限名称'"`
	url           string `gorm:"comment:'权限路径'"`
	requestMethod string `gorm:"comment:'请求方式'"`
	parentId      int64  `gorm:"comment:'父权限id'"`
	isAnonymous   bool   `gorm:"comment:'是否匿名访问'"`
}

type Role struct {
	mapper.BaseModel
	Name        string `gorm:"comment:'角色名'"`
	Description string `gorm:"comment:'角色描述'"`
	IsDisable   bool   `gorm:"comment:'是否禁用'"`
}

type UserRole struct {
	mapper.BaseModel
	UserId string `gorm:"unique;comment:'用户id'"`
	RoleId string `gorm:"unique;comment:'角色id'"`
}

type RolePermission struct {
	mapper.BaseModel
	RoleId       string `gorm:"index;comment:'角色ID'"`
	PermissionId string `gorm:"index;comment:'权限ID'"`
}
