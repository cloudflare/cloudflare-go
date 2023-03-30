package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type RequestError struct {
	Cause    error
	Request  *http.Request
	Response *http.Response
}

func (e RequestError) StatusCode() int {
	if res := e.Response; res != nil {
		return res.StatusCode
	}
	return -1
}

func (e RequestError) URL() *url.URL {
	req := e.Request
	if req == nil {
		return nil
	}
	return req.URL
}

func (e RequestError) Error() string {
	return fmt.Errorf("%s: %d: %w", e.URL(), e.StatusCode(), e.Cause).Error()
}

type APIError struct {
	request       *http.Request
	response      *http.Response
	status        int
	errorBody     interface{}
	errorBodyJSON *string
	message       string
	headers       http.Header
}

func (e APIError) Request() *http.Request {
	return e.request
}

func (e APIError) Status() int {
	return e.status
}

func (e APIError) ErrorBody() interface{} {
	return e.errorBody
}

func (e APIError) Message() string {
	return e.message
}

func (e APIError) Headers() http.Header {
	return e.headers
}

func (e APIError) errorjSON() string {
	if json := e.errorBodyJSON; json != nil {
		return *json
	}
	body := e.errorBody
	if body == nil {
		e.errorBodyJSON = &e.message
		return e.message
	}
	bytes, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		json := fmt.Sprintf("{\"error\":\"go runtime error while dumping error body: %s\"}", err.Error())
		e.errorBodyJSON = &json
		return json
	}
	json := string(bytes)
	e.errorBodyJSON = &json
	return json
}

func (e APIError) Method() string {
	if e.request == nil {
		return "<nil>"
	}
	return e.request.Method
}

func (e APIError) URL() string {
	req := e.request
	if req == nil {
		return "<nil>"
	}
	url := req.URL
	if url == nil {
		return "<nil>"
	}
	return url.String()
}

func (e APIError) Error() string {
	return fmt.Sprintf("api_error: %s %s: %d\n%s", e.Method(), e.URL(), e.status, e.errorjSON())
}

func NewAPIError(req *http.Request, res *http.Response, status int, err error, message string, headers http.Header) APIError {
	return APIError{req, res, status, err, nil, message, headers}
}

func NewAPIErrorFromResponse(req *http.Request, res *http.Response) APIError {
	message := ""

	errContent, _ := io.ReadAll(res.Body)
	message += string(errContent)

	return NewAPIError(req, res, res.StatusCode, nil, message, res.Header)
}
