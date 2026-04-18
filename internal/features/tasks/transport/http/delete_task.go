package tasks_transport_http

import (
	core_logger "gopet/internal/core/logger"
	core_http_respose "gopet/internal/core/transport/response"
	core_http_utils "gopet/internal/core/transport/utils"
	"net/http"
)

// DeleteTask   godoc
// @Summary     Delete task
// @Description Delete task in system
// @Tags        tasks
// @Param       id      path     int     true                    "ID of the task to be deleted"
// @Success     204                                              "Successfully deleted task"
// @Failure     400     {object} core_http_respose.ErorrResponse "Bad request"
// @Failure     404     {object} core_http_respose.ErorrResponse "Task not found"
// @Failure     500     {object} core_http_respose.ErorrResponse "Internal server error"
// @Router      /tasks/{id} [delete]
func (h *TasksHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	repsponseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_utils.GetIntPathValues(r, "id")
	if err != nil {
		repsponseHandler.ErorrResponse(
			err,
			"failed to get taskID path value",
		)

		return
	}

	if err := h.tasksService.DeleteTask(ctx, taskID); err != nil {
		repsponseHandler.ErorrResponse(
			err,
			"failed to delete task",
		)

		return
	}

	repsponseHandler.NoContentResponse()
}
