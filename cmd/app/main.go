package main

import (
	"hw1/internal/config"
	"hw1/internal/transport"
	"net/http"
	"time"
)

func main() {
	cfg := config.NewConfig()
	router := transport.NewRouter(cfg)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
