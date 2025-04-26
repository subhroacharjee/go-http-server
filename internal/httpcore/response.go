package httpcore

import "fmt"

type HeaderMap map[string]string

type HttpResponseWriter struct {
	statusCode    HttpStatus
	statusMessage string
	headers       HeaderMap
	body          []byte
}

func NewHttpResponseWriter() HttpResponseWriter {
	return HttpResponseWriter{
		statusCode:    StatusOK,
		statusMessage: httpStatusMessages[StatusOK],
		headers:       make(HeaderMap),
		body:          make([]byte, 0),
	}
}

func (w *HttpResponseWriter) SetStatus(httpStatus HttpStatus) {
	w.statusCode = httpStatus
	w.statusMessage = httpStatusMessages[httpStatus]
}

func (w *HttpResponseWriter) SetHeader(key string, value string) {
	w.headers[key] = value
}

func (w *HttpResponseWriter) Write(body []byte) {
	w.body = append(w.body, body...)
}

func (w HttpResponseWriter) ToResponseByte() []byte {
	separator := "\r\n"
	statusLine := []byte(fmt.Sprintf("HTTP/1.1 %d %s%s", w.statusCode, w.statusMessage, separator))

	headerLine := ""

	for key, value := range w.headers {
		headerLine += fmt.Sprintf("%s: %s%s", key, value, separator)
	}
	headerLine += separator
	headerLineBytes := []byte(headerLine)

	resp := make([]byte, 0)
	resp = append(statusLine, headerLineBytes...)

	resp = append(resp, w.body...)
	return resp
}
