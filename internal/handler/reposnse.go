package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Reposonse struct {
	Status  int
	Message string
	Data    interface{}
}

func HandleResponse(c *gin.Context, data interface{}) {
	reponse := Reposonse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}

	c.JSON(http.StatusOK, reponse)
}

func HandleError(c *gin.Context, data interface{}) {
	log.Fatalln("FAILURE:", c.Request, data)

	reponse := Reposonse{
		Status:  http.StatusBadRequest,
		Message: "Request Fail",
		Data:    data,
	}

	c.JSON(http.StatusBadRequest, reponse)
}
