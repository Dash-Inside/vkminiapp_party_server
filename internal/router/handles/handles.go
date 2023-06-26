package handles

import (
	"partyserver/internal/domain/gameuser"

	"github.com/gin-gonic/gin"
)

type handles struct {
	ur gameuser.Repository
}

type Handles interface {
	UserCreate() func(*gin.Context)
	UserRead() func(*gin.Context)
	UserUpdate() func(*gin.Context)
	UserDelete() func(*gin.Context)
}

func NewHandles(ge *gin.Engine, ur gameuser.Repository) Handles {
	return &handles{
		ur: ur,
	}
}

func ImplHandles(ge *gin.Engine, h Handles) error {
	ge.POST("/user", h.UserCreate())
	ge.GET("/user/:id", h.UserRead())
	ge.PUT("/user/:id", h.UserUpdate())
	ge.DELETE("/user/:id", h.UserDelete())

	return nil
}
