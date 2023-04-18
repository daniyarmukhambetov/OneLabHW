package transport

import (
	"context"
	"hw1/config"
	"hw1/handler"
	"log"
	"net/http"
	"time"
)

type Server struct {
	cfg  *config.Config
	Http *http.Server
	ctx  context.Context
	Conn chan int
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
	log.Println("server is starting")
	go s.ListenCtx()
	err := s.Http.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) ListenCtx() {
	log.Println("wait ctx")
	log.Println(<-s.ctx.Done(), "done")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	log.Println("gracefully shutting down")
	if err := s.Http.Shutdown(ctx); err != nil {
		log.Println(err, "err")
	}
	s.Conn <- 1
}
