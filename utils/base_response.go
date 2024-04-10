package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Messages []string    `json:"messages"`
	Data     interface{} `json:"data"`
}

type BaseWithoutMessages struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorResponseWithoutMessages(c *gin.Context, status int, message string) {
	response := BaseWithoutMessages{
		Status:  status,
		Message: message,
	}
	c.JSON(status, response)
}

func SuccessResponse(c *gin.Context, data interface{}, messages []string) {
	response := BaseResponse{
		Status:   http.StatusOK,
		Message:  "success",
		Messages: messages,
		Data:     data,
	}
	c.JSON(http.StatusOK, response)
}

func ErrorResponse(c *gin.Context, status int, message string, errs error) {
	response := BaseResponse{
		Status:   status,
		Message:  message,
		Messages: []string{errs.Error()},
	}
	c.JSON(status, response)
}
