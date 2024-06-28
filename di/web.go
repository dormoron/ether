package di

import (
	"ether/internal/handler/users"
	"ether/pkg/logger"
	"github.com/dormoron/mist"
	"github.com/dormoron/mist/middlewares/accesslog"
	"github.com/dormoron/mist/middlewares/ratelimit"
	"github.com/dormoron/mist/security"
	"github.com/redis/go-redis/v9"
	"time"
)

func InitServer(mils []mist.Middleware, userHandel users.AuthHandler, roleHandel users.RoleHandler) *mist.HTTPServer {
	server := mist.InitHTTPServer()
	// middleware
	server.Use(mils...)
	// router
	userHandel.RegisterRoutes(server)
	roleHandel.RegisterRoutes(server)
	return server
}

func InitMiddleware(redisClient redis.Cmdable, logger logger.Logger, sp security.Provider) []mist.Middleware {
	redisLimit := ratelimit.InitRedisSlidingWindowLimiter(redisClient, time.Minute, 100)
	// session provider
	security.SetDefaultProvider(sp)
	defer security.SetDefaultProvider(nil)

	return []mist.Middleware{
		accesslog.InitMiddleware().LogFunc(func(log string) {
			logger.Info(log)
		}).Build(),
		ratelimit.InitMiddlewareBuilder(redisLimit, 30).Build(),
		security.InitMiddlewareBuilder(sp, "/auth/login", "/auth/signup").Build(),
	}
}
