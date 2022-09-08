package handler

import (
	"fmt"
	"strconv"
	"worker/controller"
	"worker/server"
	"worker/utils"

	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {
	a, err := strconv.ParseInt(c.Query("a"), 10, 16)
	if err != nil {
		server.ResponseFailed(c, fmt.Sprintf("Invalid Input %s", c.Query("a")))
		return
	}
	a16 := int16(a)
	b, err := strconv.ParseInt(c.Query("b"), 10, 16)
	b16 := int16(b)
	if err != nil {
		server.ResponseFailed(c, fmt.Sprintf("Invalid Input %s", c.Query("b")))
		return
	}
	controller, err := controller.GetController(c.Param("version"))
	if err != nil {
		server.ResponseFailed(c, utils.InvalidController)
		return
	}
	datas := controller.CalculateTest(a16, b16)
	server.ResponseSucess(c, datas)
}
