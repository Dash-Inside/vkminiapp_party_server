package router

import (
	"partyserver/internal/domain/gameuser"
	"partyserver/internal/router/handles"

	"github.com/gin-gonic/gin"
)

func ProvideRouter(ur gameuser.Repository) *gin.Engine {
	engine := NewRouter()
	h := handles.NewHandles(engine, ur)
	handles.ImplHandles(engine, h)

	return engine
}
