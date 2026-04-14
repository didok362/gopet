package tasks_transport_http

import (
	core_logger "gopet/internal/core/logger"
	core_http_respose "gopet/internal/core/transport/response"
	core_http_utils "gopet/internal/core/transport/utils"
	"net/http"
)

type GetTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	repsponseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_utils.GetIntPathValues(r, "id")
	if err != nil {
		repsponseHandler.ErorrResponse(
			err,
			"failed to get taskID path value",
		)
	}

	taskDomain, err := h.tasksService.GetTask(ctx, taskID)
	if err != nil {
		repsponseHandler.ErorrResponse(
			err,
			"failed to get task",
		)

		return
	}

	resposne := GetTaskResponse(taskDTOFromDomain(taskDomain))

	repsponseHandler.JSONResponse(resposne, http.StatusOK)
}
