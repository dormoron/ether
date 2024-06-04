package model

import (
	"time"
)

type Auth struct {
	Id         int64
	UserId     int64
	Username   string
	Password   string
	CreateTime time.Time
}

type User struct {
	Id         int64
	Username   string
	WebSite    string
	Nickname   string
	Avatar     string
	Intro      string
	IsDisable  bool
	CreateTime time.Time
}
