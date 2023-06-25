package main

import (
	server "partyserver/internal/cmd"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(`assets/config.json`)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if mode := viper.GetString("GIN_MODE"); mode == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := server.RunServer(); err != nil {
		panic(err)
	}
}
