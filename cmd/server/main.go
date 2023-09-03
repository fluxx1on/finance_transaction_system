package main

import (
	"context"
	baseLog "log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fluxx1on/finance_transaction_system/internal"
	"github.com/fluxx1on/finance_transaction_system/internal/config"
	"github.com/fluxx1on/finance_transaction_system/pkg/logger"
	"golang.org/x/exp/slog"
)

func main() {
	signalCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Config
	cfg := config.Setup()
	RootDir, _ := os.Getwd()

	// Logger
	logfile, err := os.Create(RootDir + cfg.Logger.Logfile)
	if err != nil {
		baseLog.Fatalf("Logfile missed: %v", err)
	}

	log := slog.New(logger.NewColorfulHandler(
		baseLog.Default().Writer(),
		logfile,
		&slog.HandlerOptions{
			Level: logger.GetLevel(cfg.Logger.LevelInfo),
		},
	))
	slog.SetDefault(log)

	// gRPC Server starting
	server := &internal.Node{}
	server.StartUp(cfg)

	<-signalCtx.Done()

	// Shutting down
	log.Info("Server shutting down. All connection will be terminated")

	finished := make(chan struct{}, 1)
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		server.Stop()
		finished <- struct{}{}
	}()

	select {
	case <-shutdownCtx.Done():
		log.Error("Server shutdown", signalCtx.Err(), shutdownCtx.Err())
	case <-finished:
		log.Info("Successfully finished")
	}
}
