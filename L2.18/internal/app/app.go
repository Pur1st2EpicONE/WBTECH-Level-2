package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"L2.18/internal/config"
	"L2.18/internal/handler"
	"L2.18/internal/repository"
	"L2.18/internal/server"
	"L2.18/internal/service"
	"L2.18/pkg/logger"
)

type App struct {
	logger  logger.Logger
	server  server.Server
	storage repository.Storage
	ctx     context.Context
	cancel  context.CancelFunc
	wg      *sync.WaitGroup
}

func Boot() *App {

	config, err := config.Load()
	if err != nil {
		log.Fatalf("app — failed to load configs: %v", err)
	}

	logger := logger.NewLogger(config.Logger)
	server, storage := wireApp(nil, config, logger)

	ctx, cancel := newContext(logger)
	wg := new(sync.WaitGroup)

	return &App{
		logger:  logger,
		server:  server,
		storage: storage,
		ctx:     ctx,
		cancel:  cancel,
		wg:      wg,
	}

}

func wireApp(db any, config config.App, logger logger.Logger) (server.Server, repository.Storage) {
	storage := repository.NewStorage(db, config.Storage, logger)
	service := service.NewService(config.Service, storage, logger)
	handler := handler.NewHandler(service, logger)
	server := server.NewServer(config.Server, handler, logger)
	return server, storage
}

func newContext(logger logger.Logger) (context.Context, context.CancelFunc) {

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sig := <-sigCh
		logger.LogInfo("app — received signal "+sig.String()+", initiating graceful shutdown", "layer", "app")
		cancel()
	}()

	return ctx, cancel

}

func (a *App) Run() {

	a.wg.Go(func() {
		if err := a.server.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.LogFatal("server run failed", err, "layer", "app")
		}
	})

	<-a.ctx.Done()

	a.logger.LogInfo("app — shutting down...", "layer", "app")
	a.server.Shutdown()
	a.Stop()
	a.wg.Wait()

}

func (a *App) Stop() {
	a.storage.Close()
	a.logger.Close()
}
