package router

import "github.com/gin-gonic/gin"

func ProvideRouter() *gin.Engine {
	return gin.Default()
}
