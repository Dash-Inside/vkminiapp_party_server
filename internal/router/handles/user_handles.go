package handles

import (
	"partyserver/internal/domain/gameuser"
	"partyserver/internal/router/response"

	"github.com/gin-gonic/gin"
)

func (h *handles) UserCreate() func(*gin.Context) {
	return func(c *gin.Context) {
		createUserParams := gameuser.CreateUserParams{}

		if err := c.Bind(&createUserParams); err != nil {
			response.HandleError(c, err)
			return
		}

		if err := h.ur.Create(createUserParams); err != nil {
			response.HandleError(c, err)
			return
		}

		response.HandleResponse(c, createUserParams)
	}
}

func (*handles) UserDelete() func(*gin.Context) {
	panic("unimplemented")
}

func (*handles) UserRead() func(*gin.Context) {
	return func(c *gin.Context) {
		response.HandleResponse(c, map[string]string{"Nickname": "Read Test"})
	}
}

func (*handles) UserUpdate() func(*gin.Context) {
	panic("unimplemented")
}
