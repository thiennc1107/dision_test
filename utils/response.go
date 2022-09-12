package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSucess(c *gin.Context, datas interface{}) {
	c.JSON(http.StatusOK, res{
		Sucess: true,
		Status: http.StatusOK,
		Data:   datas,
	})
}

func ResponseFailed(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, res{
		Sucess:  false,
		Status:  http.StatusBadRequest,
		Message: message,
	})
}

func ResponseTimeOut(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, res{
		Sucess:  false,
		Status:  http.StatusRequestTimeout,
		Message: "Request timed out",
	})
}

func ResponseInternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, res{
		Sucess:  false,
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
	})
}

type res struct {
	Sucess  bool        `json:"sucess"`
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
