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
	server *gin.Engine
	logger interfaces.ILogger
}

var httpServerWrapper = gin.New

func (hs *HttpServer) Setup() {
	hs.server = httpServerWrapper()
	// hs.server.Use(hs.logger.GetHandleFunc())
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

func (hs HttpServer) Run() error {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	return hs.server.Run(fmt.Sprintf("%s:%s", host, port))
}

func NewHttpServer(logger interfaces.ILogger) IHttpServer {
	return &HttpServer{
		logger: logger,
	}
}
