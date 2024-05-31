package model

import "time"

type Auth struct {
	Id         int64
	Username   string
	Password   string
	IsDisable  bool
	CreateTime time.Time
}
