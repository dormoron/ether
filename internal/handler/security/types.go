package security

import (
	"github.com/dormoron/mist"
	"github.com/golang-jwt/jwt/v5"
)

type Handler interface {
	SetLoginToken(ctx *mist.Context, uid int64) error
	SetJWTToken(ctx *mist.Context, uid int64, ssid string) error
	ClearToken(ctx *mist.Context) error
	CheckSession(ctx *mist.Context, ssid string) error
	ExtractToken(ctx *mist.Context) string
}

type RefreshClaims struct {
	Uid  int64
	Ssid string
	jwt.RegisteredClaims
}

type UserClaims struct {
	jwt.RegisteredClaims
	// 声明你自己的要放进去 token 里面的数据
	Id   int64
	Ssid string
	// 自己随便加
	UserAgent string
}
