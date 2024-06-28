package security

import (
	"errors"
	"fmt"
	"github.com/dormoron/mist"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

var (
	AccessTokenKey  = []byte("7sR7cvUbGb4uy1q4r7GIdy5Ue0sOPdQP")
	RefreshTokenKey = []byte("7sR7cvUbGb4uy1q4r7GIdy5Ue0sOPdQE")
)

type JwtHandler struct {
	cmd redis.Cmdable
}

func NewJwtHandler(cmd redis.Cmdable) Handler {
	return &JwtHandler{
		cmd: cmd,
	}
}

func (h *JwtHandler) SetLoginToken(ctx *mist.Context, uid int64) error {
	ssid := uuid.New().String()
	err := h.SetJWTToken(ctx, uid, ssid)
	if err != nil {
		return err
	}
	err = h.setRefreshToken(ctx, uid, ssid)
	return err
}

func (h *JwtHandler) SetJWTToken(ctx *mist.Context, uid int64, ssid string) error {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
		Id:        uid,
		Ssid:      ssid,
		UserAgent: ctx.Request.UserAgent(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString(AccessTokenKey)
	if err != nil {
		return err
	}
	ctx.Header("x-access-security", tokenStr)
	return nil
}

func (h *JwtHandler) ClearToken(ctx *mist.Context) error {
	ctx.Header("x-jwt-token", "")
	ctx.Header("x-refresh-token", "")

	claims := ctx.MustGet("claims").(*UserClaims)
	return h.cmd.Set(ctx, fmt.Sprintf("users:ssid:%s", claims.Ssid),
		"", time.Hour*24*7).Err()
}

func (h *JwtHandler) CheckSession(ctx *mist.Context, ssid string) error {
	val, err := h.cmd.Exists(ctx, fmt.Sprintf("users:ssid:%s", ssid)).Result()
	switch {
	case errors.Is(err, redis.Nil):
		return nil
	case err == nil:
		if val == 0 {
			return nil
		}
		return errors.New("session 已经无效了")
	default:
		return err
	}
}

func (h *JwtHandler) ExtractToken(ctx *mist.Context) string {
	tokenHeader := ctx.Request.Header.Get("Authorization")
	segs := strings.Split(tokenHeader, " ")
	if len(segs) != 2 {
		return ""
	}
	return segs[1]
}

func (h *JwtHandler) setRefreshToken(ctx *mist.Context, uid int64, ssid string) error {
	claims := RefreshClaims{
		Ssid: ssid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
		Uid: uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString(RefreshTokenKey)
	if err != nil {
		return err
	}
	ctx.Header("x-refresh-security", tokenStr)
	return nil
}
