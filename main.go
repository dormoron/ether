package main

import (
	"ether/integration"
	"github.com/spf13/viper"
)

func main() {
	initViperReader()
	server := integration.InitWebServer()
	server.Start(":8080")

}

func initViperReader() {
	viper.SetConfigFile("config/dev_test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
