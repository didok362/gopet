package users_transport_http

import (
	"gopet/internal/core/domain"
	core_logger "gopet/internal/core/logger"
	core_http_request "gopet/internal/core/transport/request"
	core_http_respose "gopet/internal/core/transport/response"
	"net/http"
)

type CreateUserRequset struct {
	FullName    string  `json:"full_name"    validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,starwith=+"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke CreateUser")
	var request CreateUserRequset
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErorrResponse(err, "decode and validate error")
	}
	rw.WriteHeader(http.StatusOK)

	userDomain := domainFromDTO(request)

	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErorrResponse(err, "failed to create user")
	}
}

func domainFromDTO(dto CreateUserRequset) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}
