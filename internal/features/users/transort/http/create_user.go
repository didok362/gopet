package users_transport_http

import (
	"gopet/internal/core/domain"
	core_logger "gopet/internal/core/logger"
	core_http_request "gopet/internal/core/transport/request"
	core_http_respose "gopet/internal/core/transport/response"
	"net/http"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"    validate:"required,min=3,max=100"               example:"Ivan Ivanon"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+" example:"+38097123804"`
}

type CreateUserResponse UserDTOResponse

// CreateUser godoc
// @Summary     Create User
// @Description Create new user in system
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       request body CreateUserRequest true          "CreateUser request body"
// @Success     201 {object} CreateUserResponse              "Successfully created user"
// @Failure     400 {object} core_http_respose.ErorrResponse "Bad request"
// @Failure     500 {object} core_http_respose.ErorrResponse "Internal server error"
// @Router      /users [post]
func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke CreateUser")
	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErorrResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain := domainFromDTO(request)

	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErorrResponse(err, "failed to create user")
		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)

}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}
