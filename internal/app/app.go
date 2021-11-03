// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"id-maker/config"
	v1 "id-maker/internal/controller/http/v1"
	"id-maker/internal/controller/rpc"
	"id-maker/internal/usecase"
	"id-maker/internal/usecase/repo"
	"id-maker/pkg/grpcserver"
	"id-maker/pkg/httpserver"
	"id-maker/pkg/logger"
	"id-maker/pkg/mysql"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	mysql, err := mysql.New(
		cfg.MySQL.URL,
		mysql.MaxIdleConns(cfg.MySQL.MaxIdleConns),
		mysql.MaxOpenConns(cfg.MySQL.MaxOpenConns),
	)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - mysql.New: %w", err))
	}
	defer mysql.Close()

	// Use case
	segmentUseCase := usecase.New(
		repo.New(mysql),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, segmentUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// gRPC Server
	grpcServer := grpcserver.New(grpcserver.Port(cfg.GRPC.Port))
	rpc.NewRouter(segmentUseCase, l)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-grpcServer.Notify():
		l.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	grpcServer.Shutdown()
}
