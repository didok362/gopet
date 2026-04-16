package main

import (
	"context"
	"fmt"
	core_config "gopet/internal/core/config"
	core_logger "gopet/internal/core/logger"
	core_postgres_pool "gopet/internal/core/repository/postgres/pool"
	core_http_middleware "gopet/internal/core/transport/middleware"
	core_http_server "gopet/internal/core/transport/server"
	statistics_postgres_repository "gopet/internal/features/statistics/repository"
	statistics_service "gopet/internal/features/statistics/service"
	statistics_transport_http "gopet/internal/features/statistics/transport/http"
	tasks_postgres_repository "gopet/internal/features/tasks/repositroy"
	tasks_service "gopet/internal/features/tasks/service"
	tasks_transport_http "gopet/internal/features/tasks/transport/http"
	users_postgres_repository "gopet/internal/features/users/repository/postgres"
	users_service "gopet/internal/features/users/service"
	users_transport_http "gopet/internal/features/users/transort/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	cfg := core_config.NewConfigMust()
	time.Local = cfg.TimeZone

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("falied to init logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("app time zone:", zap.Any("zone", time.Local))

	logger.Debug("init postgres pool")
	pool, err := core_postgres_pool.NewConnectionPool(
		ctx,
		core_postgres_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("Failed to connect postgress pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("init feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("init feature", zap.String("feature", "tasks"))
	tasksRepository := tasks_postgres_repository.NewTaskRepository(pool)
	tasksService := tasks_service.NewTaskService(tasksRepository)
	tasksTransportHTTP := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	logger.Debug("init statistics", zap.String("feature", "statistics"))
	statisticsRepository := statistics_postgres_repository.NewStatisticsRepository(pool)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransportHTTP := statistics_transport_http.NewStatisticsHTTPHandler(statisticsService)

	logger.Debug("init of http srever")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegiseterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouter.RegiseterRoutes(tasksTransportHTTP.Routes()...)
	apiVersionRouter.RegiseterRoutes(statisticsTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
