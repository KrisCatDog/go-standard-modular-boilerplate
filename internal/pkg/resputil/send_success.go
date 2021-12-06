package resputil

import (
	"github.com/gin-gonic/gin"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
)

type baseSuccessResponse struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendSuccess(c *gin.Context, httpCode int, msg string, data interface{}) {
	c.JSON(httpCode, &baseSuccessResponse{
		Code:    api.CodeSuccess,
		Message: msg,
		Data:    data,
	})
}
