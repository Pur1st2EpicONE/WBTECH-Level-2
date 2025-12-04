// Package app defines the main application structure and lifecycle management.
//
// It handles application initialization, context and signal management, server startup,
// graceful shutdown, and resource cleanup. The App struct encapsulates all components
// required to run the calendar service, including logger, server, storage, and context.
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

// App represents the main application instance, managing its components and lifecycle.
type App struct {
	logger  logger.Logger      // Structured logger used throughout the application for info, warning, error, and debug logs
	server  server.Server      // HTTP server instance that handles incoming requests
	storage repository.Storage // Persistent storage layer for events and application data
	ctx     context.Context    // Context used for cancellation and graceful shutdown
	cancel  context.CancelFunc // Function to cancel the application context and trigger shutdown
	wg      *sync.WaitGroup    // WaitGroup to synchronize goroutines during server run and shutdown
}

// Boot initializes the application and returns an App instance.
//
// This function performs the following tasks:
//  1. Loads configuration from files or environment variables.
//  2. Initializes the structured logger.
//  3. Wires together storage, service, handler, and HTTP server components.
//  4. Sets up a cancellable context that listens to OS signals for graceful shutdown.
//  5. Creates a wait group for managing goroutines.
//
// If any critical error occurs during initialization (e.g., configuration load failure),
// the function logs the error and terminates the application.
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

// wireApp initializes repository, service, handler, and server components.
//
// It returns the fully configured HTTP server and storage instance.
// This function allows optional dependency injection for the database (db parameter).
func wireApp(db any, config config.App, logger logger.Logger) (server.Server, repository.Storage) {
	storage := repository.NewStorage(db, config.Storage, logger)
	service := service.NewService(config.Service, storage, logger)
	handler := handler.NewHandler(service, logger)
	server := server.NewServer(config.Server, handler, logger)
	return server, storage
}

// newContext creates a cancellable context and listens to OS signals for graceful shutdown.
//
// The function sets up a goroutine that waits for SIGINT or SIGTERM signals.
// When a signal is received, it logs the event and cancels the context, which triggers
// shutdown procedures in the App.Run method.
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

// Run starts the HTTP server and waits for shutdown signals.
//
// It performs the following:
//  1. Starts the HTTP server in a separate goroutine managed by the wait group.
//  2. Waits for the application context to be cancelled (e.g., OS signal received).
//  3. Logs the shutdown initiation.
//  4. Calls server.Shutdown() to gracefully stop accepting new requests.
//  5. Calls App.Stop() to release resources such as storage and logger.
//  6. Waits for all goroutines to finish using the wait group.
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

// Stop releases all application resources.
//
// It closes the storage and logger to ensure proper cleanup before exiting.
func (a *App) Stop() {
	a.storage.Close()
	a.logger.Close()
}
