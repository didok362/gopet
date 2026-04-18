package users_transport_http

import (
	"fmt"
	"gopet/internal/core/domain"
	core_logger "gopet/internal/core/logger"
	core_http_request "gopet/internal/core/transport/request"
	core_http_respose "gopet/internal/core/transport/response"
	core_http_types "gopet/internal/core/transport/types"
	core_http_utils "gopet/internal/core/transport/utils"
	"net/http"
	"strings"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nulladble[string] `json:"full_name"    swaggertype:"string" example:"Iv1n Invaov"`
	PhoneNumber core_http_types.Nulladble[string] `json:"phone_number" swaggertype:"string" example:"+38097223804"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Value == nil {
		if r.FullName.Set {
			return fmt.Errorf("FullName cant be empty")
		}

		FullNameLen := len([]rune(*r.FullName.Value))

		if FullNameLen < 3 || FullNameLen > 100 {
			return fmt.Errorf("full name should be more than 3 and less them 100")
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			PhoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if PhoneNumberLen < 10 || PhoneNumberLen > 15 {
				return fmt.Errorf("PhoneNuber must be more then 10 and less then 15")
			}

			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("'PhoneNumber' must strat with +")
			}
		}
	}

	return nil
}

type PatchUserResponse UserDTOResponse

// PatchUser    godoc
// @Summary     Patch User
// @Description ### Three-state logic:
// @Description 1. **Field is not provided:** nothing to be done
// @Description 2. **Field is provided(+12390481234):** creating new
// @Description 3. **Field is provided(null):** set to null
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       id path int true "ID of the user to be patched"
// @Param       request body PatchUserRequest true "PatchUser request body"
// @Success     200 {object} PatchUserResponse               "Successfully patched user"
// @Failure     400 {object} core_http_respose.ErorrResponse "Bad request"
// @Failure     404 {object} core_http_respose.ErorrResponse "Not found"
// @Failure     409 {object} core_http_respose.ErorrResponse "Conflict"
// @Failure     500 {object} core_http_respose.ErorrResponse "Internal server error"
// @Router      /users/{id} [patch]
func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	var request PatchUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	userPatch := userPatchFromRequset(request)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to patch",
		)

		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func userPatchFromRequset(request PatchUserRequest) domain.UserPatch {
	return domain.UserPatch{
		FullName:    request.FullName.ToDomain(),
		PhoneNumber: request.PhoneNumber.ToDomain(),
	}
}
