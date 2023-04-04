package main

import (
	"context"
	"errors"
	"hw1/internal/config"
	"hw1/internal/logger"
	"hw1/internal/transport"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.NewConfig()
	router := transport.NewRouter(cfg)
	srv := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	shutdownError := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.Logger().Println("shutting down server")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		shutdownError <- srv.Shutdown(ctx)
	}()
	logger.Logger().Println("starting server")
	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return
	}
	err = <-shutdownError
	if err != nil {
		return
	}
	logger.Logger().Println("server is stopped")
}
