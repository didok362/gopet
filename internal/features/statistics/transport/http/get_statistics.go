package statistics_transport_http

import (
	"fmt"
	"gopet/internal/core/domain"
	core_logger "gopet/internal/core/logger"
	core_http_respose "gopet/internal/core/transport/response"
	core_http_utils "gopet/internal/core/transport/utils"
	"net/http"
	"time"
)

type GetStatisticsResponse struct {
	TasksCreated             int      `json:"tasks_created"`
	TasksCompleted           int      `json:"tasks_completed"`
	TasksCompletedRate       *float64 `json:"tasks_completed_rate"`
	TasksAverageCompleteTime *string  `json:"tasks_average_complte_time"`
}

// GetStatistics godoc
// @Summary      Get Statistics
// @Description  Get statistics with optional filters such as: user if/time gap
// @Tags         statistics
// @Produce      json
// @Param        user_id      query       int       false     "user id filtration"
// @Param        from         query       string    false     "time scinse starting to analyse statistics"
// @Param        to           query       string    false     "time till analyse statistics "
// @Success      200 {object} GetStatisticsResponse           "Successfully got statistics"
// @Failure      400 {object} core_http_respose.ErorrResponse "Bad request"
// @Failure      500 {object} core_http_respose.ErorrResponse "Internal server error"
// @Router       /statistics [get]
func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_respose.NewHTTPResponseHandler(log, rw)

	UserID, from, to, err := getUserIDFromToQueryParams(r)
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to get param from query",
		)

		return
	}

	statistics, err := h.statisticsService.GetStatistics(ctx, UserID, from, to)
	if err != nil {
		responseHandler.ErorrResponse(
			err,
			"failed to get statistics",
		)

		return
	}

	response := toDTOFromDomain(statistics)

	responseHandler.JSONResponse(response, http.StatusOK)

}

func toDTOFromDomain(statistics domain.Statistics) GetStatisticsResponse {
	var avgTime *string

	if statistics.TasksAverageCompleteTime != nil {
		duration := statistics.TasksAverageCompleteTime.String()
		avgTime = &duration
	}

	return GetStatisticsResponse{
		TasksCreated:             statistics.TasksCreated,
		TasksCompleted:           statistics.TasksCompleted,
		TasksCompletedRate:       statistics.TasksCompletedRate,
		TasksAverageCompleteTime: avgTime,
	}
}

func getUserIDFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	userID, err := core_http_utils.GetIntQueryParam(r, "user_id")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get user_id query param: %w", err)
	}

	from, err := core_http_utils.GetDateQueryParam(r, "from")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get from query param: %w", err)
	}

	to, err := core_http_utils.GetDateQueryParam(r, "to")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get to query param: %w", err)
	}

	return userID, from, to, nil
}
