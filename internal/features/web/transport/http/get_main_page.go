package web_http_transport

import (
	core_logger "gopet/internal/core/logger"
	core_http_respose "gopet/internal/core/transport/response"
	"net/http"
)

func (h *WebHTTPHandler) GetMainPage(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)

	html, err := h.webService.GetMainPage()
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to get index html for main page",
		)

		return
	}

	responseHandler.HTMLResponse(html)
}
