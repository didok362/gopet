package users_transport_http

import (
	"fmt"
	core_logger "gopet/internal/core/logger"
	core_http_respose "gopet/internal/core/transport/response"
	core_http_utils "gopet/internal/core/transport/utils"
	"net/http"
)

type GetUsersResponse []UserDTOResponse

// GetUsers     godoc
// @Summary     Get Users
// @Description Get existing users in system with optional pagination
// @Tags        users
// @Produce		json
// @Param       limit  query int false 						 "Pagination limit parameter"
// @Param       offset query int false 						 "Pagination offset parameter"
// @Success     200 {object} GetUsersResponse                "Successfully retrieved users"
// @Failure     400 {object} core_http_respose.ErorrResponse "Bad request"
// @Failure     500 {object} core_http_respose.ErorrResponse "Internal server error"
// @Router      /users [get]
func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to get limit/offset query params",
		)
		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to get users",
		)
		return
	}

	response := GetUsersResponse(usersDTOFromDoamins(userDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get limit query param: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get offset query param: %w", err)
	}

	return limit, offset, nil
}
