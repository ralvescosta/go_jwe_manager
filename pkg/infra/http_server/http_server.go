package http_server

import (
	"context"
	"jwemanager/pkg/app/errors"
	"jwemanager/pkg/app/interfaces"
	"net/http"
	"time"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type IHttpServer interface {
	Default()
	RegistreRoute(method string, path string, handlers ...gin.HandlerFunc) error
	Setup()
	Run() error
}

type HttpServer struct {
	env      interfaces.IEnvironments
	logger   interfaces.ILogger
	router   *gin.Engine
	server   *http.Server
	shotdown chan bool
}

var httpServerWrapper = gin.New

func (pst *HttpServer) Default() {
	pst.router = httpServerWrapper()
	pst.router.Use(GinLogger(pst.logger))
	pst.router.SetTrustedProxies(nil)
}

func (hs HttpServer) RegistreRoute(method string, path string, handlers ...gin.HandlerFunc) error {
	switch method {
	case "POST":
		hs.router.POST(path, handlers...)
	case "GET":
		hs.router.GET(path, handlers...)
	case "PUT":
		hs.router.PUT(path, handlers...)
	case "PATCH":
		hs.router.PATCH(path, handlers...)
	case "DELETE":
		hs.router.DELETE(path, handlers...)
	default:
		return errors.NewInternalError("http method not allowed")
	}
	return nil
}

func (pst *HttpServer) Setup() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("%s:%s", host, port)

	pst.server = &http.Server{
		Addr:    addr,
		Handler: pst.router,
	}

	go pst.gracefullShutdown()
	// if pst.env.GO_ENV() != pst.env.PROD_ENV() {
	// 	certPath := os.Getenv("TLS_CERT_PATH")
	// 	keyPath := os.Getenv("TLS_KEY_PATH")
	// 	return pst.router.RunTLS(addr, certPath, keyPath)
	// }
}

func (pst HttpServer) Run() error {
	err := pst.server.ListenAndServe()
	if err != nil {
		errors.NewInternalError(err.Error())
	}

	return nil
}

func (pst HttpServer) gracefullShutdown() {
	<-pst.shotdown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pst.server.Shutdown(ctx); err != nil {
		pst.logger.Error("[HttpServer::GracefullShutdown] - could'ent shutdown properly")
		return
	}
}

func NewHttpServer(environments interfaces.IEnvironments, logger interfaces.ILogger, shotdown chan bool) IHttpServer {
	return &HttpServer{
		env:      environments,
		logger:   logger,
		shotdown: shotdown,
	}
}
