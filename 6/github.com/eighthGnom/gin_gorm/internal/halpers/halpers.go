package halpers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Message struct {
	StatusCode int         `json:"status_code"`
	Meta       interface{} `json:"meta"`
	Data       interface{} `json:"data"`
}

func RespondJSON(ctx *gin.Context, status_code int, data interface{}) {
	log.Println("status code: ", status_code)
	message := Message{
		StatusCode: status_code,
		Data:       data,
	}
	ctx.JSON(status_code, message)
}
