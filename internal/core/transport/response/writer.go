package core_http_respose

import "net/http"

var (
	StatusCodeUninitiaized = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     StatusCodeUninitiaized,
	}
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.statusCode = statusCode

}

func (rw *ResponseWriter) GetStatusCodeOrPanic() int {
	if rw.statusCode == StatusCodeUninitiaized {
		panic("no status code set")
	}
	return rw.statusCode
}
