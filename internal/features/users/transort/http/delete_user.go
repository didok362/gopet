package users_transport_http

import (
	core_logger "gopet/internal/core/logger"
	core_http_respose "gopet/internal/core/transport/response"
	core_http_utils "gopet/internal/core/transport/utils"
	"net/http"
)

func (h *UsersHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to get userID from path",
		)

		return
	}

	if err := h.usersService.DeleteUser(ctx, userID); err != nil {
		responseHandler.ErorrResponse(
			err,
			"faield to delete user",
		)

		return
	}

	responseHandler.NoContentResponse()

}
