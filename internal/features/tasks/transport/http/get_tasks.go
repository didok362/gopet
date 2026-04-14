package tasks_transport_http

import (
	"fmt"
	core_logger "gopet/internal/core/logger"
	core_http_respose "gopet/internal/core/transport/response"
	core_http_utils "gopet/internal/core/transport/utils"
	"net/http"
)

type GetTasksResponse []TaskDTOResponse

func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)
	userID, limit, offset, err := getUseridLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to get userID/lim/offest from query params",
		)

		return
	}

	tasksDomains, err := h.tasksService.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to get tasks from query params",
		)

		return
	}

	response := GetTasksResponse(tasksDTOsFromDomains(tasksDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUseridLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	userID, err := core_http_utils.GetIntQueryParam(r, "user_id")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get userid query param: %w", err)
	}

	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get limit query param: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get offset query param: %w", err)
	}

	return userID, limit, offset, nil
}
