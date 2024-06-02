package main

import (
	"github.com/dormoron/mist"
	"github.com/dormoron/mist/middlewares/accesslog"
	"github.com/dormoron/mist/middlewares/opentelemetry"
	"github.com/dormoron/mist/middlewares/prometheus"
	"github.com/dormoron/mist/middlewares/recovery"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	server := initServer()
	server.Start(":8080")

}

func initServer() *mist.HTTPServer {
	server := mist.InitHTTPServer()
	server.Use(
		opentelemetry.InitMiddlewareBuilder().Build(),
		accesslog.InitMiddleware().LogFunc(
			func(log string) {
				zap.L().Info(log)
			}).Build(),
		recovery.InitMiddlewareBuilder(
			http.StatusInternalServerError,
			[]byte("internal server error")).SetLogFunc(func(ctx *mist.Context, err any) {
			zap.L().Error("internal server error", zap.Any("panic", err),
				zap.String("route", ctx.MatchedRoute))
		}).Build(),
		prometheus.InitMiddlewareBuilder(
			"ether", "web", "ether", "ether 的 web 统计").Build())
	return server
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("username:password@0@tcp(ip:3306)/ether"))
	if err != nil {
		panic(err)
	}
	return db
}
