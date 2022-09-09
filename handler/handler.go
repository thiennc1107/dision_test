package handler

import (
	"strconv"
	"worker/controller"
	"worker/server"

	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {
	a, err := strconv.ParseInt(c.Query("a"), 10, 16)
	if err != nil {
		server.ResponseFailed(c, server.InvalidInput(c.Query("a")))
		return
	}
	a16 := int16(a)
	b, err := strconv.ParseInt(c.Query("b"), 10, 16)
	b16 := int16(b)
	if err != nil {
		server.ResponseFailed(c, server.InvalidInput(c.Query("b")))
		return
	}
	controller, err := controller.GetController(c.Param("version"))
	if err != nil {
		server.ResponseFailed(c, err.Error())
		return
	}
	datas, err := controller.CalculateTest(a16, b16)
	if err != nil {
		server.ResponseFailed(c, err.Error())
		return
	}
	err = datas.CheckInvalidResult()
	if err != nil {
		server.ResponseFailed(c, err.Error())
		return
	}
	server.ResponseSucess(c, datas)
}
