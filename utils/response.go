package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSucess(c *gin.Context, datas interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"sucess": true,
		"status": http.StatusOK,
		"data":   datas,
	})
}

func ResponseFailed(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"sucess":  false,
		"status":  http.StatusBadRequest,
		"message": message,
	})
}

func ResponseInternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"sucess":  false,
		"status":  http.StatusInternalServerError,
		"message": "internal server error",
	})
}
