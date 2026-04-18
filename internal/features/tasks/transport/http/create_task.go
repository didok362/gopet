package tasks_transport_http

import (
	"gopet/internal/core/domain"
	core_logger "gopet/internal/core/logger"
	core_http_request "gopet/internal/core/transport/request"
	core_http_respose "gopet/internal/core/transport/response"
	"net/http"
)

type CreateTaskRequest struct {
	Title        string  `json:"title"          validate:"required,min=1,max=100"    example:"make my math homework"`
	Description  *string `json:"description"    validate:"omitempty,min=1,max=1000"  example:"i gotta have a math lesson tomorrow so i need to do my hw"`
	AuthorUserID int     `json:"author_user_id" validate:"required"                  example:"123"`
}

type CreateTaskResponse TaskDTOResponse

// CreateTask   godoc
// @Summary     Create task
// @Description Create new task in system
// @Tags        tasks
// @Accept      json
// @Produce     json
// @Param       request body     CreateTaskRequest true          "CreateTask request body"
// @Success     200     {object} CreateTaskResponse              "Successfully created task"
// @Failure     400     {object} core_http_respose.ErorrResponse "Bad request"
// @Failure     404     {object} core_http_respose.ErorrResponse "Author not found"
// @Failure     500     {object} core_http_respose.ErorrResponse "Internal server error"
// @Router      /tasks [post]
func (h *TasksHTTPHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)

	var request CreateTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErorrResponse(
			err,
			"failede to decode and vlaidate HTTP request",
		)

		return
	}

	taskDomain := domain.NewTaskUninitialized(
		request.Title,
		request.Description,
		request.AuthorUserID,
	)

	taskDomain, err := h.tasksService.CreateTask(ctx, taskDomain)
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to create task",
		)

		return
	}

	response := CreateTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}
