package di

import (
	"ether/internal/domain/repository/mapper"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	_ = mapper.InitTable(db)
	return db
}
