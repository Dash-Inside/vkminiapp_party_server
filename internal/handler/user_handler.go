package handler

import (
	"partyserver/internal/domain/gameuser"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateUser(c *gin.Context) {
	createUserParams := gameuser.CreateUserParams{}

	if err := c.Bind(&createUserParams); err != nil {
		HandleError(c, err)
		return
	}

	if err := h.userReposiory.Create(createUserParams); err != nil {
		HandleError(c, err)
		return
	}

	HandleResponse(c, createUserParams)
}

func (*handler) DeleteUser(*gin.Context) {
	panic("unimplemented")
}

func (*handler) ReadUser(*gin.Context) {
	panic("unimplemented")
}

func (*handler) UpdateUser(*gin.Context) {
	panic("unimplemented")
}
