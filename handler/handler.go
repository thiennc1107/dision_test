package handler

import (
	"crypto/tls"
	"net/http"
	"strconv"
	"worker/models"
	"worker/utils"

	"github.com/gin-gonic/gin"
)

type IController interface {
	CalculateTest(a, b int16) (models.Data, error)
	Log(msg string)
	IsDebug() bool
}
type ApiSevice struct {
	server     *http.Server
	controller IController
}

func (s *ApiSevice) Serve(cert, key string) error {
	return s.server.ListenAndServeTLS(cert, key)
}

func (s *ApiSevice) TestHandler(c *gin.Context) {
	a, err := strconv.ParseInt(c.Query("a"), 10, 16)
	s.controller.Log("This is a test log")
	if err != nil {
		utils.ResponseFailed(c, utils.InvalidInput(c.Query("a")))
		return
	}
	a16 := int16(a)
	b, err := strconv.ParseInt(c.Query("b"), 10, 16)
	b16 := int16(b)
	if err != nil {
		utils.ResponseFailed(c, utils.InvalidInput(c.Query("b")))
		return
	}
	datas, err := s.controller.CalculateTest(a16, b16)
	if err != nil {
		utils.ResponseFailed(c, err.Error())
		return
	}
	err = datas.CheckInvalidResult()
	if err != nil {
		utils.ResponseFailed(c, err.Error())
		return
	}
	utils.ResponseSucess(c, datas)
}

func (s *ApiSevice) loadServer() *ApiSevice {
	if !s.controller.IsDebug() {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	server := http.Server{
		Addr:    ":1234",
		Handler: r,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	}
	s.server = &server

	r.GET("/api/:version/test", s.TestHandler)
	return s
}

func NewAPiService(controller IController) *ApiSevice {
	apiService := new(ApiSevice)
	apiService.controller = controller
	apiService.loadServer()
	return apiService
}
