package server

import (
	"context"
	"partyserver/internal/database"
	"partyserver/internal/handler"
	"partyserver/internal/logger"
	"partyserver/internal/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func loop(lifecycle fx.Lifecycle, router *gin.Engine) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			address := viper.GetString(`ADDRESS`)
			port := viper.GetString(`PORT`)

			if err := router.Run(address + ":" + port); err != nil {
				return err
			}

			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})
}

func RunServer() error {
	app := fx.New(
		fx.Provide(logger.NewLogger),
		fx.Provide(database.ProvideDatabase),
		fx.Provide(router.ProvideRouter),
		fx.Provide(handler.NewHandler),
		fx.Invoke(loop),
	)

	if err := app.Err(); err != nil {
		return err
	}

	app.Run()

	return nil
}
