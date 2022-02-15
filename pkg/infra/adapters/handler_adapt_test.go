package adapters

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"jwemanager/pkg/app/interfaces"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/infra/logger"

	"github.com/gin-gonic/gin"
)

func Test_Should_Exec_Handler_Successfully(t *testing.T) {
	sut := makeSut()

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 1 {
		t.Error("should called handler once")
	}
}

func Test_Should_Exec_Handler_With_Body_Error(t *testing.T) {
	sut := makeSut()

	readAllBody = func(r io.Reader) ([]byte, error) {
		return []byte{}, errors.New("Error")
	}

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 0 {
		t.Error("Shouldn't call handler when body is unformatted")
	}
}

func makeSut() sutReturn {
	handlerCalledTimes := 0
	handlerMock := func(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
		handlerCalledTimes++
		return httpServer.HttpResponse{}
	}

	loggerMock := logger.NewLoggerSpy()

	req := &http.Request{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte(nil))),
		Header: http.Header{
			"op": []string{"op"},
		},
	}

	sut := HandlerAdapt(handlerMock, loggerMock)

	contextMock, _ := gin.CreateTestContext(httptest.NewRecorder())
	contextMock.Params = []gin.Param{{Key: "key", Value: "value"}}
	contextMock.Request = req

	return sutReturn{
		handlerMock:        handlerMock,
		handlerCalledTimes: &handlerCalledTimes,
		loggerMock:         loggerMock,
		request:            req,
		adapt:              sut,
		ctx:                contextMock,
	}
}

type sutReturn struct {
	adapt              gin.HandlerFunc
	loggerMock         interfaces.ILogger
	handlerCalledTimes *int
	handlerMock        func(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
	request            *http.Request
	ctx                *gin.Context
}
