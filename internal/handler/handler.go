package handler

import "github.com/gin-gonic/gin"

type handler struct {
	engine *gin.Engine
}

type Handler interface {
	init() error
}

func NewGameHandler(ge *gin.Engine) Handler {
	return &handler{
		engine: ge,
	}
}

func (h *handler) init() error {
	h.engine.POST("/user")
	h.engine.GET("/user/:id")
	h.engine.PUT("/user/:id")
	h.engine.DELETE("/user/:id")

	return nil
}
