package middleware

import (
	"ether/internal/handler/security"
	"github.com/dormoron/mist"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

var JwtKey = []byte("7sR7cvUbGb4uy1q4r7GIdy5Ue0sOPdQP")

type LoginMiddlewareBuilder struct {
	paths []string
	security.Handler
}

func NewLoginMiddlewareBuilder(handler security.Handler) *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{
		Handler: handler,
	}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginMiddlewareBuilder) Build() mist.Middleware {
	return func(next mist.HandleFunc) mist.HandleFunc {
		return func(ctx *mist.Context) {
			for _, path := range l.paths {
				if ctx.Request.URL.Path == path {
					next(ctx)
					return
				}
			}

			tokenStr := l.ExtractToken(ctx)
			claims := &security.UserClaims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return JwtKey, nil
			})
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			if token == nil || !token.Valid {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			if claims.UserAgent != ctx.Request.UserAgent() {
				// todo 加监控，移动token
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			err = l.CheckSession(ctx, claims.Ssid)
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			next(ctx)
		}
	}
}
