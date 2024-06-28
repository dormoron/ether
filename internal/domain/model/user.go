package model

import (
	"time"
)

type Auth struct {
	Id         uint
	UserId     uint
	Username   string
	Password   string
	CreateTime time.Time
}

type User struct {
	Id         uint
	Username   string
	WebSite    string
	Nickname   string
	Avatar     string
	Intro      string
	IsDisable  bool
	CreateTime time.Time
}

type Role struct {
	Id          uint
	Name        string
	Description string
	IsDefault   bool
	IsDisable   bool
	UpdateTime  time.Time
	CreateTime  time.Time
}
