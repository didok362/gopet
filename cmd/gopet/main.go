package main

import (
	"context"
	"fmt"
	core_logger "gopet/internal/core/logger"
	core_http_middleware "gopet/internal/core/transport/middleware"
	core_http_server "gopet/internal/core/transport/server"
	users_transport_http "gopet/internal/features/users/transort/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
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

	logger.Debug("star of apl!")
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(nil)
	userRoutes := usersTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegiseterRoutes(userRoutes...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
