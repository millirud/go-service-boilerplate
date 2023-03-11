package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/millirud/go-service-boilerplate/config"
	"github.com/millirud/go-service-boilerplate/internal/controller/http/middlewares"
	"github.com/millirud/go-service-boilerplate/internal/controller/http/probes"
	"github.com/millirud/go-service-boilerplate/pkg/httpserver"
	"golang.org/x/net/context"
)

func Run(cfg *config.Config) {
	var err error
	ctx, cancel := context.WithCancel(context.Background())

	handler := gin.New()
	handler.Use(middlewares.NewLogger(), gin.Recovery())

	handler.GET("/healthz", probes.NewLivenessProbe())

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	handler.GET("/healthz/ready", probes.NewReadinessProbe(ctx))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		cancel()
		fmt.Printf("app - Run - signal: %s /n", s.String())
	case err = <-httpServer.Notify():
		cancel()
		fmt.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		fmt.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
