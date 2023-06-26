package handler

import (
	"log"
	"partyserver/internal/domain/gameuser"

	"github.com/gin-gonic/gin"
)

type handler struct {
	// engine        *gin.Engine
	userReposiory gameuser.Repository
}

type Handler interface {
	CreateUser(*gin.Context)
	ReadUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

func NewHandler(ge *gin.Engine) Handler {
	handler := &handler{
		// engine: ge,
	}

	log.Println("Handler init")

	ge.POST("/user", handler.CreateUser)
	ge.GET("/user/:id", handler.ReadUser)
	ge.PUT("/user/:id", handler.UpdateUser)
	ge.DELETE("/user/:id", handler.DeleteUser)

	return handler
}
