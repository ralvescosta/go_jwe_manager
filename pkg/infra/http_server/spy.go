package httpServer

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type HttpServerSpy struct {
	mock.Mock
}

func (pst HttpServerSpy) Default() {}
func (pst HttpServerSpy) RegistreRoute(method string, path string, handlers ...gin.HandlerFunc) error {
	args := pst.Called(method, path, handlers)

	return args.Error(0)
}
func (pst HttpServerSpy) Setup() {

}
func (pst HttpServerSpy) Run() error {
	args := pst.Called()

	return args.Error(0)
}

func NewHttpServerSpy() *HttpServerSpy {
	return new(HttpServerSpy)
}
