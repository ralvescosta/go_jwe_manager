package http_server

import (
	"jwemanager/pkg/app/interfaces"

	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type IHttpServer interface {
	Setup()
	RegistreRoute(method string, path string, handlers ...gin.HandlerFunc) error
	Run() error
}

type HttpServer struct {
	env    interfaces.IEnvironments
	logger interfaces.ILogger
	server *gin.Engine
}

var httpServerWrapper = gin.New

func (pst *HttpServer) Setup() {
	pst.server = httpServerWrapper()
	pst.server.Use(GinLogger(pst.logger))
	pst.server.SetTrustedProxies(nil)
}

func (hs HttpServer) RegistreRoute(method string, path string, handlers ...gin.HandlerFunc) error {
	switch method {
	case "POST":
		hs.server.POST(path, handlers...)
	case "GET":
		hs.server.GET(path, handlers...)
	case "PUT":
		hs.server.PUT(path, handlers...)
	case "PATCH":
		hs.server.PATCH(path, handlers...)
	case "DELETE":
		hs.server.DELETE(path, handlers...)
	default:
		return errors.New("http method not allowed")
	}
	return nil
}

func (pst HttpServer) Run() error {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	url := fmt.Sprintf("%s:%s", host, port)

	if pst.env.GO_ENV() != pst.env.PROD_ENV() {
		certPath := os.Getenv("TLS_CERT_PATH")
		keyPath := os.Getenv("TLS_KEY_PATH")
		return pst.server.RunTLS(url, certPath, keyPath)
	}

	return pst.server.Run(url)
}

func NewHttpServer(environments interfaces.IEnvironments, logger interfaces.ILogger) IHttpServer {
	return &HttpServer{
		env:    environments,
		logger: logger,
	}
}
