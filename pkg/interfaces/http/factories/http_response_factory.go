package factories

import (
	"net/http"

	"jwemanager/pkg/app/errors"
	httpserver "jwemanager/pkg/infra/http_server"
	vm "jwemanager/pkg/interfaces/http/view_models"
)

type HttpResponseFactory struct{}

func (HttpResponseFactory) Ok(body interface{}, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 200,
		Body:       body,
		Headers:    headers,
	}
}

func (HttpResponseFactory) Created(body interface{}, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 201,
		Body:       body,
		Headers:    headers,
	}
}

func (HttpResponseFactory) NoContent(headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 204,
		Headers:    headers,
	}
}

func (HttpResponseFactory) BadRequest(msg string, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 400,
		Body: vm.ErrorMessage{
			StatusCode: 400,
			Message:    msg,
		},
		Headers: headers,
	}
}

func (HttpResponseFactory) Unauthorized(msg string, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 401,
		Body: vm.ErrorMessage{
			StatusCode: 401,
			Message:    msg,
		},
		Headers: headers,
	}
}

func (HttpResponseFactory) Forbidden(msg string, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 403,
		Body: vm.ErrorMessage{
			StatusCode: 403,
			Message:    msg,
		},
		Headers: headers,
	}
}

func (HttpResponseFactory) NotFound(msg string, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 404,
		Body: vm.ErrorMessage{
			StatusCode: 404,
			Message:    msg,
		},
		Headers: headers,
	}
}

func (HttpResponseFactory) Conflict(msg string, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 409,
		Body: vm.ErrorMessage{
			StatusCode: 409,
			Message:    msg,
		},
		Headers: headers,
	}
}

func (HttpResponseFactory) InternalServerError(msg string, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: 500,
		Body: vm.ErrorMessage{
			StatusCode: 500,
			Message:    msg,
		},
		Headers: headers,
	}
}

func (HttpResponseFactory) GenericResponse(statusCode int, body interface{}, headers http.Header) httpserver.HttpResponse {
	return httpserver.HttpResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    headers,
	}
}

func (pst HttpResponseFactory) ErrorResponseMapper(err error, headers http.Header) httpserver.HttpResponse {
	switch err.(type) {
	case errors.NotFoundError:
		return pst.NotFound(err.Error(), headers)
	case errors.ConflictError:
		return pst.Conflict(err.Error(), headers)
	default:
		return pst.InternalServerError(err.Error(), headers)
	}
}

func NewHttpResponseFactory() HttpResponseFactory {
	return HttpResponseFactory{}
}
