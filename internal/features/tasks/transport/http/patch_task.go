package tasks_transport_http

import (
	"fmt"
	"gopet/internal/core/domain"
	core_logger "gopet/internal/core/logger"
	core_http_request "gopet/internal/core/transport/request"
	core_http_respose "gopet/internal/core/transport/response"
	core_http_types "gopet/internal/core/transport/types"
	core_http_utils "gopet/internal/core/transport/utils"
	"net/http"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nulladble[string] `json:"title"       swaggertype:"string"  example:"Make HW"`
	Description core_http_types.Nulladble[string] `json:"description" swaggertype:"string"  example:"i need to make hw"`
	Completed   core_http_types.Nulladble[bool]   `json:"completed"   swaggertype:"boolean" example:"false"`
}

func (r *PatchTaskRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("Titlte cant be null")
		}

		titleLen := len([]rune(*r.Title.Value))
		if titleLen < 1 || titleLen > 100 {
			return fmt.Errorf("title must be between 1 and 100")
		}
	}

	if r.Description.Set {
		if r.Description.Value != nil {
			descriptionLen := len([]rune(*r.Description.Value))
			if descriptionLen < 1 || descriptionLen > 1000 {
				return fmt.Errorf("description must be more than 1 and less than 1000")
			}
		}
	}

	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("Completed cant be null")
		}
	}

	return nil
}

type PatchUserResponse TaskDTOResponse

// PatchTask    godoc
// @Summary     Patch task
// @Description Get tasks in system
// @Description ### Three-state logic:
// @Description 1. **Field is not provided:** nothing to be done
// @Description 2. **Field is provided("go for a walk"):** redact
// @Description 3. **Field is provided(null):** set to null
// @Tags        tasks
// @Produce		json
// @Accept 		json
// @Param       id      path     int              true           "ID of the task to be patched"
// @Param 		request body     PatchTaskRequest true           "request body of PatchTask"
// @Success     200     {object} PatchUserResponse               "Successfully patched task"
// @Failure     400     {object} core_http_respose.ErorrResponse "Bad request"
// @Failure     404     {object} core_http_respose.ErorrResponse "Task not found"
// @Failure     409     {object} core_http_respose.ErorrResponse "Conflict"
// @Failure     500     {object} core_http_respose.ErorrResponse "Internal server error"
// @Router      /tasks/{id} [patch]
func (h *TasksHTTPHandler) PatchTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_utils.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to get id",
		)

		return
	}

	var request PatchTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to decode and validate http request",
		)

		return
	}

	taskPatch := TaskPatchFromRequest(request)

	taskDomain, err := h.tasksService.PatchTask(ctx, taskID, taskPatch)
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to patch task",
		)

		return
	}

	response := PatchUserResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func TaskPatchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.NewTaskPatch(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Completed.ToDomain(),
	)
}
