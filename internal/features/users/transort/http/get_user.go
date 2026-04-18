package users_transport_http

import (
	core_logger "gopet/internal/core/logger"
	core_http_respose "gopet/internal/core/transport/response"
	core_http_utils "gopet/internal/core/transport/utils"
	"net/http"
)

type GetUserResponse UserDTOResponse

// GetUser      godoc
// @Summary     Get User
// @Description Get existing user in system
// @Tags        users
// @Produce		json
// @Param       id path int true                             "ID of the user to be retrieved"
// @Success     200 {object} GetUserResponse                 "Successfully retrieved user"
// @Failure     400 {object} core_http_respose.ErorrResponse "Bad request"
// @Failure     404 {object} core_http_respose.ErorrResponse "Not found"
// @Failure     500 {object} core_http_respose.ErorrResponse "Internal server error"
// @Router      /users/{id} [get]
func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to get user id from path",
		)

		return
	}

	user, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to get user from userID",
		)
		return
	}

	response := GetUserResponse(userDTOFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}
