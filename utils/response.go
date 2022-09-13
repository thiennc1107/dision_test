package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSucess(c *gin.Context, datas interface{}) {
	c.JSON(http.StatusOK, Res{
		Success: true,
		Status:  http.StatusOK,
		Data:    datas,
	})
}

func ResponseFailed(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Res{
		Success: false,
		Status:  http.StatusBadRequest,
		Message: message,
	})
}

func ResponseTimeOut(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, Res{
		Success: false,
		Status:  http.StatusRequestTimeout,
		Message: "Request timed out",
	})
}

func ResponseInternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Res{
		Success: false,
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
	})
}

type Res struct {
	Success bool        `json:"sucess"`
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
