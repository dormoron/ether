package users

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id         int64          `gorm:"primaryKey,autoIncrement"`
	Email      sql.NullString `gorm:"unique;comment:'email'"`
	WebSite    string         `gorm:"comment:'个人网站'"`
	Nickname   string         `gorm:"comment:'用户昵称'"`
	Avatar     string         `gorm:"comment:'用户头像'"`
	Intro      string         `gorm:"comment:'用户简介'"`
	IsDisable  bool           `gorm:"comment:'是否禁用'"`
	CreateTime int64          `gorm:"comment:'创建时间'"`
	UpdateTime int64          `gorm:"comment:'修改时间'"`
	DeleteAt   gorm.DeletedAt `gorm:"index;column:'软删除'"`
}

type Auth struct {
	Id            int64          `gorm:"primaryKey,autoIncrement"`
	UserId        sql.NullString `gorm:"comment:'用户id'"`
	Username      sql.NullString `gorm:"unique;comment:'用户名'"`
	Password      string         `gorm:"comment:'密码'"`
	LoginType     string         `gorm:"comment:'登录类型'"`
	IpAddress     string         `gorm:"comment:'用户登录ip'"`
	IpSource      string         `gorm:"comment:'ip来源'"`
	LastLoginTime time.Time      `gorm:"comment:'上次登录时间'"`
	CreateTime    int64          `gorm:"comment:'创建时间'"`
	UpdateTime    int64          `gorm:"comment:'修改时间'"`
	DeleteAt      gorm.DeletedAt `gorm:"index;column:'软删除'"`
}

type Role struct {
	Id          int64          `gorm:"primaryKey,autoIncrement"`
	Name        string         `gorm:"comment:'角色名'"`
	Description string         `gorm:"comment:'角色描述'"`
	IsDisable   bool           `gorm:"comment:'是否禁用'"`
	CreateTime  int64          `gorm:"comment:'创建时间'"`
	UpdateTime  int64          `gorm:"comment:'修改时间'"`
	DeleteAt    gorm.DeletedAt `gorm:"index;column:'软删除'"`
}

type UserRole struct {
	Id         int64          `gorm:"primaryKey,autoIncrement"`
	UserId     string         `gorm:"not null;comment:'用户id'"`
	RoleId     string         `gorm:"not null;comment:'角色id'"`
	CreateTime int64          `gorm:"comment:'创建时间'"`
	UpdateTime int64          `gorm:"comment:'修改时间'"`
	DeleteAt   gorm.DeletedAt `gorm:"index;column:'软删除'"`
}
