package core_http_respose

import (
	"encoding/json"
	"errors"
	"fmt"
	core_errors "gopet/internal/core/errors"
	core_logger "gopet/internal/core/logger"
	"net/http"

	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponseHandler(log *core_logger.Logger, rw http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		rw:  rw,
	}
}

func (h *HTTPResponseHandler) JSONResponse(
	responseBody any,
	statusCode int,
) {
	h.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(h.rw).Encode(responseBody); err != nil {
		h.log.Error("write http response", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) NoContentResponse() {
	h.rw.WriteHeader(http.StatusNoContent)
}

func (h *HTTPResponseHandler) HTMLResponse(html []byte) {
	h.rw.WriteHeader(http.StatusOK)

	h.rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := h.rw.Write(html); err != nil {
		h.log.Error("write HTML HTTP respose", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) ErorrResponse(err error, msg string) {
	var (
		statusCode int
		logfunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgumnet):
		statusCode = http.StatusBadRequest
		logfunc = h.log.Warn
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logfunc = h.log.Debug
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logfunc = h.log.Debug
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logfunc = h.log.Error
	default:
		statusCode = http.StatusInternalServerError
		logfunc = h.log.Error
	}

	logfunc(msg, zap.Error(err))

	h.errorResponse(
		statusCode,
		err,
		msg,
	)
}

func (h *HTTPResponseHandler) PanicReponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic %v", p)

	h.log.Error(msg, zap.Error(err))
	h.errorResponse(
		statusCode,
		err,
		msg,
	)
}

func (h *HTTPResponseHandler) errorResponse(
	statusCode int,
	err error,
	msg string,
) {
	h.rw.WriteHeader(statusCode)

	response := ErorrResponse{
		Error:   err.Error(),
		Message: msg,
	}

	h.JSONResponse(
		response,
		statusCode,
	)
}
