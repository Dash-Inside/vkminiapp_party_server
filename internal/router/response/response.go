package response

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int
	Message string
	Data    interface{}
}

func HandleResponse(c *gin.Context, data interface{}) {
	response := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}

func HandleError(c *gin.Context, data interface{}) {
	log.Fatalln("FAILURE:", c.Request, data)

	response := Response{
		Status:  http.StatusBadRequest,
		Message: "Request Fail",
		Data:    data,
	}

	c.JSON(http.StatusBadRequest, response)
}
