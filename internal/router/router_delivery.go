package router

import (
	"partyserver/internal/handler"

	"github.com/gin-gonic/gin"
)

func ProvideRouter() *gin.Engine {
	engine := gin.Default()
	handler.NewHandler(engine)

	return engine
}
