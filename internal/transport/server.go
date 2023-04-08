package transport

import (
	"context"
	"hw1/internal/config"
	"hw1/internal/handler"
	"hw1/internal/logger"
	"net/http"
	"time"
)

type Server struct {
	cfg  *config.Config
	Http *http.Server
	ctx  context.Context
}

func NewServer(cfg *config.Config, ctx context.Context, handler *handler.Handler) (*Server, error) {
	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: NewRouter(handler),
	}
	return &Server{
		cfg:  cfg,
		ctx:  ctx,
		Http: server,
	}, nil
}

func (s *Server) Run() error {
	logger.Logger().Println("server is starting")
	go s.ListenCtx()
	err := s.Http.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) ListenCtx() {
	logger.Logger().Println("wait ctx")
	logger.Logger().Println(<-s.ctx.Done(), "done")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.Logger().Println("gracefully shutting down")
	if err := s.Http.Shutdown(ctx); err != nil {
		logger.Logger().Println(err)
	}
}
