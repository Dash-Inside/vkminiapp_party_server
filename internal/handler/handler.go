package handler

import "github.com/gin-gonic/gin"

type handler struct {
	engine *gin.Engine
}

type Handler interface {
	CreateUser(*gin.Context)
	ReadUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

func NewGameHandler(ge *gin.Engine) Handler {
	handler := &handler{
		engine: ge,
	}

	ge.POST("/user", handler.CreateUser)
	ge.GET("/user/:id", handler.ReadUser)
	ge.PUT("/user/:id", handler.UpdateUser)
	ge.DELETE("/user/:id", handler.DeleteUser)

	return handler
}
