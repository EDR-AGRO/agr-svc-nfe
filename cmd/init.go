package cmd

import (
	"br/com/agr/nfe/infrastructure/apm"
	"br/com/agr/nfe/infrastructure/logger"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	apmt := apm.StartTransaction(ctx, "StartApp", "StartApp")
	defer apmt.EndTransaction()
	defer func() {
		apm.Flush()
		time.Sleep(5 * time.Second) // Wait for APM to flush
	}()

	builder := NewBuilder()
	// defer builder.Close(apmt)

	buildDirector := NewAppBuilder(builder)

	errChan, err := buildDirector.Build(apmt)
	if err != nil {
		logger.Panicf(apmt.Ctx, "Building error: %v", err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		logger.Infof(apmt.Ctx, "Shutdown signal received")
		logger.Infof(apmt.Ctx, "Gracefully shutting down the application...")
	case err = <-errChan:
		if err != nil {
			logger.Panicf(apmt.Ctx, "Error starting app: %v", err.Error())
		} else {
			logger.Infof(apmt.Ctx, "App stopped gracefully")
		}
	}

	logger.Infof(apmt.Ctx, "Application exiting")
}
