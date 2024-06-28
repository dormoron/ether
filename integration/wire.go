//go:build wireinject

package integration

import (
	"ether/di"
	"ether/internal/domain/repository/cache"
	userMapper "ether/internal/domain/repository/mapper/users"
	userRepository "ether/internal/domain/repository/users"
	userService "ether/internal/domain/service/users"
	userHandel "ether/internal/handler/users"
	"github.com/dormoron/mist"
	"github.com/google/wire"
)

func InitWebServer() *mist.HTTPServer {
	wire.Build(
		// base
		di.InitDB, di.InitRedis, di.InitLogger,

		// mapper
		userMapper.NewUserMapper,
		userMapper.NewUserAuthMapper,
		userMapper.NewRoleMapper,
		userMapper.NewUserRoleMapper,

		// cache
		cache.NewUserCache,

		// repository
		userRepository.NewAuthRepository,
		userRepository.NewUserRepository,
		userRepository.NewRoleRepository,
		userRepository.NewUserRoleRepository,

		// service
		userService.NewUserService,
		userService.NewRoleService,

		// handler
		userHandel.NewAuthHandler,
		userHandel.NewRoleHandler,

		// access
		di.InitAccess,

		// session
		di.InitSessionProvider,

		// web server
		di.InitServer, di.InitMiddleware,
	)
	return new(mist.HTTPServer)
}
