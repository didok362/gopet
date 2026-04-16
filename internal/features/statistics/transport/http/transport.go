package statistics_transport_http

import (
	"context"
	"gopet/internal/core/domain"
	core_http_server "gopet/internal/core/transport/server"
	"net/http"
	"time"
)

type StatisticsHTTPHandler struct {
	statisticsService StatisticsService
}

type StatisticsService interface {
	GetStatistics(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(
	statisticsService StatisticsService,
) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		statisticsService: statisticsService,
	}
}

func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
