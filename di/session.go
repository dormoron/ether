package di

import (
	"github.com/dormoron/mist/security"
	"github.com/dormoron/mist/security/redisess"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitSessionProvider(client redis.Cmdable) security.Provider {
	return redisess.InitSessionProvider(client, viper.GetString("jwt.key"))
}
